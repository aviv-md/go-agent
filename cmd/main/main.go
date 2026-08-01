package main

import (
	"context"
	"explorations/agents/internal/agent"
	"explorations/agents/internal/infra/config"
	"explorations/agents/internal/model"

	"github.com/charmbracelet/log"
)

func main() {

	env := config.Load()

	ctx := context.Background()

	m := model.NewOpenAIProvider("~deepseek/deepseek-v4-flash-latest", env.APIKey, env.BaseURL)
	a := agent.NewAgent(
		m,
		10,
	)

	query := "Is there anybody home?"

	resp, err := a.Run(ctx, query)
	if err != nil {
		log.Fatal(err)
	}

	log.Info(resp)
}
