package model

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
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

func convertMessageToOpenAIInput(message Message) (responses.ResponseInputItemUnionParam, error) {

	r, err := convertRoleToOpenAIRole(message.Role)
	if err != nil {
		return responses.ResponseInputItemUnionParam{}, err
	}

	result := responses.ResponseInputItemParamOfMessage(message.Content, r)
	return result, nil
}

func convertOpenAIResponseToModelResponse(resp *responses.Response) Response {

	return Response{
		Content: resp.OutputText(),
	}
}

// Prompt implements [Model].
func (o *OpenAIProvider) Prompt(ctx context.Context, input []Message) (Response, error) {
	i := []responses.ResponseInputItemUnionParam{}

	for _, message := range input {
		converted, err := convertMessageToOpenAIInput(message)
		if err != nil {
			return Response{}, err
		}

		i = append(i, converted)
	}

	response, err := o.client.Responses.New(ctx, responses.ResponseNewParams{
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: i,
		},
		Model: o.model,
	})

	if err != nil {
		return Response{}, err
	}

	return convertOpenAIResponseToModelResponse(response), nil
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
