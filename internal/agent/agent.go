package agent

import (
	"context"
	"explorations/agents/internal/model"
	"explorations/agents/internal/tools"
	"fmt"
	"strings"
	"unicode"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const (
	defaultMaxIterations = 10
)

type agent struct {
	MaxIterations uint8
	Model         model.Model
	Registry      tools.Registry
	SystemMessage string
}

func (a *agent) Run(ctx context.Context, input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input must be provided")
	}

	// track MaxIterations
	var iterations uint8 = 0

	// Create first msg
	messages := []model.Message{
		model.NewSystemMessage(a.SystemMessage),
		model.NewUserMessage(input),
	}

	registryTools := a.Registry.List()
	modelTools := model.Tools{}

	for _, t := range registryTools {
		modelTools = append(modelTools, model.NewAvailableTool(t.Name(), t.Description(), t.Parameters()))
	}

	for {
		assistantMessage, err := a.Model.Prompt(ctx, messages, modelTools)

		if err != nil {
			return "", err
		}

		messages = append(messages, assistantMessage)

		iterations++

		if len(assistantMessage.ToolCalls()) == 0 {
			break
		}

		thought := strings.TrimRightFunc(assistantMessage.Content(), unicode.IsSpace)
		log.Printf("%s %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#00fefc")).Render("[Thought]"), thought)
		if iterations >= a.MaxIterations {
			return messages[len(messages)-1].Content(), fmt.Errorf("max iterations reached")
		}

		for _, t := range assistantMessage.ToolCalls() {
			// execute the tool and retrieve observation
			log.Printf("%s %s %v", lipgloss.NewStyle().Foreground(lipgloss.Color("#CFFF04")).Render("[Action]"), t.Name(), t.Args())
			content, err := a.Registry.Execute(t.Name(), t.Args())

			if err != nil {
				return "", err
			}

			// append tool call and observation to messages
			obs := model.NewToolMessage(content, t.ID())
			log.Printf("%s %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#2CFF05")).Render("[Observation]"), obs.Content())
			messages = append(messages, obs)
		}
	}
	resp := messages[len(messages)-1].Content()
	log.Printf("%s %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Render("[Respond]"), resp)
	return resp, nil
}

func NewAgent(m model.Model, maxIterations uint8, registry tools.Registry, systemMessage string) *agent {

	if maxIterations == 0 {
		log.Warn("MaxIterations must be greater than 0\n Moving to default")
		maxIterations = defaultMaxIterations
	}

	if m == nil {
		panic("Model must be provided")
	}

	return &agent{
		Model:         m,
		Registry:      registry,
		MaxIterations: maxIterations,
		SystemMessage: systemMessage,
	}
}
