package agent

import (
	"context"
	"explorations/agents/internal/model"
	"fmt"

	"github.com/charmbracelet/log"
)

const (
	defaultMaxIterations = 10
)

type agent struct {
	MaxIterations uint8
	Model         model.Model
}

func (a *agent) Run(ctx context.Context, input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input must be provided")
	}

	// track MaxIterations
	var iterations uint8 = 0

	// Create first msg
	messages := []model.Message{
		model.Message{
			Role:    model.RoleUser,
			Content: input,
		},
	}

	for {
		resp, err := a.Model.Prompt(ctx, messages)

		if err != nil {
			return "", err
		}

		messages = append(messages, model.Message{
			Role:    model.RoleAssistant,
			Content: resp.Content,
		})

		iterations++

		if len(resp.ToolCalls) == 0 {
			break
		}

		if iterations == a.MaxIterations {
			return messages[len(messages)-1].Content, fmt.Errorf("max iterations reached")
		}

		for _, t := range resp.ToolCalls {
			log.Print("Tool call", t)
			messages = append(messages, model.Message{
				Role:    model.RoleTool,
				Content: "Tool calls are not implemented",
			})
		}
	}

	return messages[len(messages)-1].Content, nil
}

func NewAgent(m model.Model, maxIterations uint8) *agent {

	if maxIterations == 0 {
		log.Warn("MaxIterations must be greater than 0\n Moving to default")
		maxIterations = defaultMaxIterations
	}

	if m == nil {
		log.Fatal("Model must be provided")
	}

	return &agent{
		Model:         m,
		MaxIterations: maxIterations,
	}
}
