package model

type (
	Response struct {
		Content   string
		ToolCalls []any // currently a stub.
	}

	Message struct {
		Role    Role
		Content string
	}

	Model interface {
		Prompt(input []Message) (Response, error)
	}

	Role = string
)

const (
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleSystem    = "system"
	RoleTool      = "tool"
)
