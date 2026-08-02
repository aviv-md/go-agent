package agent

import (
	"context"
	"explorations/agents/internal/model"
	"fmt"
	"testing"
)

// Implementing a fake model.

type MockModel struct {
	messages []model.AssistantMessage
	current  int
}

// Prompt implements [model.Model].
func (m *MockModel) Prompt(ctx context.Context, input []model.Message) (model.AssistantMessage, error) {
	if m.current >= len(m.messages) {
		return nil, fmt.Errorf("no more responses")
	}

	resp := m.messages[m.current]
	m.current = m.current + 1

	return resp, nil
}

func TestRun_NoToolCalls_ReturnsContent(t *testing.T) {
	// Arrange
	mockModel := MockModel{
		messages: []model.AssistantMessage{
			model.NewAssistantMessage("Hello, world!", []model.ToolCall{}),
		},
	}

	agent := NewAgent(&mockModel, uint8(10))

	r, err := agent.Run(t.Context(), "Hey there")
	if err != nil {
		t.Fatal(err)
	}

	if mockModel.current != 1 {
		t.Fatalf("expected 1 call, got %d", mockModel.current)
	}

	if r != "Hello, world!" {
		t.Fatalf("expected 'Hello, world!', got '%s'", r)
	}
}

func TestRun_ToolCalls_LowBudget(t *testing.T) {
	mockModel := MockModel{
		messages: []model.AssistantMessage{
			model.NewAssistantMessage(
				`Searching for file "test.md"`,
				[]model.ToolCall{
					model.NewToolCall(
						"1",
						"find_file",
						map[string]any{
							"file_name": "test.md",
						},
					)},
			),
			model.NewAssistantMessage(
				`Reading file "test.md"`,
				[]model.ToolCall{
					model.NewToolCall(
						"2",
						"read_file",
						map[string]any{
							"file_name": "test.md",
						},
					)},
			),
		},
	}

	agent := NewAgent(&mockModel, 2)
	r, err := agent.Run(t.Context(), "Look for test.md file on my computer and read it")

	if err == nil {
		t.Fatal("Expected run to return err, instead err is nil")
	}

	if r != `Reading file "test.md"` {
		t.Fatalf("Expected string: %s, instead it returned empty string", `Reading file "test.md"`)
	}

	if mockModel.current != 2 {
		t.Fatalf("Expected 2 calls, got %d", mockModel.current)
	}
}
