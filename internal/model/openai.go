package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/openai/openai-go/v3/responses"
)

type OpenAIProvider struct {
	model  string
	client openai.Client
}

// standard compile-time interface assertion.
var _ Model = (*OpenAIProvider)(nil)

func convertRoleToOpenAIRole(role Role) (responses.EasyInputMessageRole, error) {
	switch role {
	case RoleAssistant:
		return responses.EasyInputMessageRoleAssistant, nil
	case RoleSystem:
		return responses.EasyInputMessageRoleSystem, nil
	case RoleUser:
		return responses.EasyInputMessageRoleUser, nil
	}

	return "", fmt.Errorf(`role "%s" cannot be converted to OpenAI`, role)
}

func convertMessageToOpenAIInput(message Message) ([]responses.ResponseInputItemUnionParam, error) {

	var result []responses.ResponseInputItemUnionParam

	switch message.Role() {
	case RoleSystem, RoleUser:
		r, _ := convertRoleToOpenAIRole(message.Role())
		result = append(result, responses.ResponseInputItemParamOfMessage(message.Content(), r))
	case RoleAssistant:
		r, _ := convertRoleToOpenAIRole(message.Role())
		am, _ := message.(AssistantMessage) // I expect it to always work bc of the switch

		if am.Content() != "" {
			result = append(result, responses.ResponseInputItemParamOfMessage(message.Content(), r))
		}

		toolCalls := am.ToolCalls()

		for _, toolCall := range toolCalls {
			args, err := json.Marshal(toolCall.Args())
			if err != nil {
				return nil, err
			}
			result = append(result, responses.ResponseInputItemParamOfFunctionCall(string(args), toolCall.id, toolCall.name))
		}
	case RoleTool:
		// Tool output bro.
		tm := message.(ToolMessage)

		if tm.CallID() == "" {
			return nil, errors.New("tool message must have a call id")
		}

		result = append(result, responses.ResponseInputItemParamOfFunctionCallOutput(tm.CallID(), tm.Content()))
	}

	return result, nil
}

func convertOpenAIResponseToAssistantMessage(resp *responses.Response) (AssistantMessage, error) {

	// Convert toolcalls.
	o := resp.Output
	content := resp.OutputText()
	toolCalls := []ToolCall{}

	for _, item := range o {
		if item.Type != "function_call" {
			continue
		}

		// Unmarshal JSON to map[string]any
		var args map[string]any
		err := json.Unmarshal([]byte(item.Arguments.OfString), &args)
		if err != nil {
			return nil, err
		}

		toolCalls = append(toolCalls, NewToolCall(
			item.CallID,
			item.Name,
			args,
		))
	}

	return NewAssistantMessage(content, toolCalls), nil
}

func convertToolsToOpenAITools(tool AvailableTool) responses.ToolUnionParam {
	t := responses.ToolParamOfFunction(tool.Name(), tool.Parameters(), false)

	t.OfFunction.Description = param.NewOpt(tool.Description())

	return t
}

// Prompt implements [Model].
func (o *OpenAIProvider) Prompt(ctx context.Context, input []Message, tools Tools) (AssistantMessage, error) {
	convertedInputs := []responses.ResponseInputItemUnionParam{}
	convertedTools := []responses.ToolUnionParam{}

	for _, message := range input {

		converted, err := convertMessageToOpenAIInput(message)
		if err != nil {
			return nil, err
		}

		convertedInputs = append(convertedInputs, converted...)
	}

	for _, tool := range tools {
		converted := convertToolsToOpenAITools(tool)
		convertedTools = append(convertedTools, converted)
	}

	response, err := o.client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: convertedInputs,
		},
		Tools: convertedTools,
		Model: o.model,
	})

	if err != nil {
		return nil, err
	}

	// Convert response to AssistantMessage.
	result, err := convertOpenAIResponseToAssistantMessage(response)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewOpenAIProvider(model string, apiKey string, baseURL string) *OpenAIProvider {

	opts := []option.RequestOption{
		option.WithAPIKey(apiKey),
		option.WithBaseURL(baseURL),
	}

	oai := openai.NewClient(opts...)

	p := &OpenAIProvider{
		model,
		oai,
	}

	return p
}
