package model

import "context"

type (
	Tools = []AvailableTool

	Model interface {
		Prompt(ctx context.Context, input []Message, tools Tools) (AssistantMessage, error)
	}

	Role string
)

const (
	RoleUser      = Role("user")
	RoleAssistant = Role("assistant")
	RoleSystem    = Role("system")
	RoleTool      = Role("tool")
)
