package agent

import (
	"context"
	"explorations/agents/internal/model"
	"fmt"
	"testing"
)

// Implementing a fake model.

type MockModel struct {
	responses []model.Response
	current   int
}

// Prompt implements [model.Model].
func (m *MockModel) Prompt(ctx context.Context, input []model.Message) (model.Response, error) {
	if m.current >= len(m.responses) {
		return model.Response{}, fmt.Errorf("no more responses")
	}

	resp := m.responses[m.current]
	m.current = m.current + 1

	return resp, nil
}

func TestRun_NoToolCalls_ReturnsContent(t *testing.T) {
	// Arrange
	mockModel := MockModel{
		responses: []model.Response{
			{
				Content: "Hello, world!",
			},
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
		responses: []model.Response{
			{
				Content: `Searching for file "test.md"`,
				ToolCalls: []any{
					"find_file",
				},
			},
			{
				Content: `Reading file "test.md"`,
				ToolCalls: []any{
					"read_file",
				},
			},
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
