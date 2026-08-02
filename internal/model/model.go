package model

import "context"

type (
	Model interface {
		Prompt(ctx context.Context, input []Message) (AssistantMessage, error)
	}

	Role string
)

const (
	RoleUser      = Role("user")
	RoleAssistant = Role("assistant")
	RoleSystem    = Role("system")
	RoleTool      = Role("tool")
)
