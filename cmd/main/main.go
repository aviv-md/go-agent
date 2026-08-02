package main

import (
	"context"
	"explorations/agents/internal/agent"
	"explorations/agents/internal/infra/config"
	"explorations/agents/internal/model"
	"explorations/agents/internal/tools"

	"github.com/charmbracelet/log"
)

func main() {

	env := config.Load()

	ctx := context.Background()

	systemMessage := `
	You are a home management agent.
	When a tool is needed provide a brief visible "Thought/status"
	Call exactly one tool per turn.
	The tool call itself is an "Action"

	a tool call result is an "Observation"
	After receiving an "Observation", you can either:
	- Provide the final answer to the user
	- Call a tool again

	when no tool is needed return a message without a tool call

	So your life cycle is:
	Thought => Action => Observation => Respond
	`

	r := tools.NewRegistry()
	temprature := tools.NewRoomTemperatureTool()

	r.Register(temprature)

	m := model.NewOpenAIProvider("~deepseek/deepseek-v4-flash-latest", env.APIKey, env.BaseURL)
	a := agent.NewAgent(
		m,
		10,
		*r,
		systemMessage,
	)

	query := "Do you know the temprature in my bedroom?"

	_, err := a.Run(ctx, query)
	if err != nil {
		log.Fatal(err)
	}
}
