package main

import (
	"context"
	"explorations/agents/internal/infra/config"

	"github.com/charmbracelet/log"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

func main() {

	env := config.Load()

	// ai.Provider
	// In order to instantiate this one we need to instantiate
	// the openai client

	oai := openai.NewClient(
		option.WithBaseURL(env.AI.BaseURL),
		option.WithAPIKey(env.AI.APIKey),
	)

	// Let's set an example prompt!
	q := "Anybody home?"

	ctx := context.Background()
	resp, err := oai.Responses.New(ctx,
		responses.ResponseNewParams{
			Input: responses.ResponseNewParamsInputUnion{
				OfString: openai.String(q),
			},
			Model: openai.ChatModelGPT5Nano,
		},
	)

	if err != nil {
		log.Error(err)
	}

	log.Info(resp.OutputText())
}
