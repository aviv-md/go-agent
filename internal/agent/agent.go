package agent

import (
	"context"
	"explorations/agents/internal/model"

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
}

func NewAgent(m model.Model, maxIterations uint8) *agent {

	if maxIterations == 0 {
		log.Warn("MaxIterations must be greater than 0\n Moving to default")
		maxIterations = defaultMaxIterations
	}

	if m == nil {
		log.Fatal("Model must be provided")
	}

	// Return a val and not a ptr as I don't manage agent
	// State yet. it shouldn't change for now.
	return &agent{
		Model:         m,
		MaxIterations: maxIterations,
	}
}
