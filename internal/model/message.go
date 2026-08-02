package model

type (
	ToolCall struct {
		id   string
		name string
		args map[string]any
	}

	message struct {
		role      Role
		content   string
		toolCalls []ToolCall
		callID    string
	}

	Message interface {
		Role() Role
		Content() string
	}

	UserMessage   = Message
	SystemMessage = Message

	AssistantMessage interface {
		Message
		ToolCalls() []ToolCall
	}

	ToolMessage interface {
		Message
		CallID() string
	}
)

func NewUserMessage(content string) UserMessage {
	return message{
		role:    RoleUser,
		content: content,
	}
}

func NewSystemMessage(content string) SystemMessage {
	return message{
		role:    RoleSystem,
		content: content,
	}
}

func NewAssistantMessage(content string, toolCalls []ToolCall) AssistantMessage {
	return message{
		role:      RoleAssistant,
		content:   content,
		toolCalls: toolCalls,
	}
}

func NewToolMessage(content string, callID string) ToolMessage {
	return message{
		role:    RoleTool,
		content: content,
		callID:  callID,
	}
}

func NewToolCall(id string, name string, args map[string]any) ToolCall {
	return ToolCall{
		id:   id,
		name: name,
		args: args,
	}
}

// Message Methods
// Value bc a message is never changing
func (m message) Role() Role {
	return m.role
}

func (m message) Content() string {
	return m.content
}

func (m message) ToolCalls() []ToolCall {
	return m.toolCalls
}

func (m message) CallID() string {
	return m.callID
}

// Tool Methods
func (t ToolCall) ID() string {
	return t.id
}

func (t ToolCall) Name() string {
	return t.name
}

func (t ToolCall) Args() map[string]any {
	return t.args
}
