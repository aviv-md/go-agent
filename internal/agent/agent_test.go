package agent

import (
	"context"
	"explorations/agents/internal/model"
	"explorations/agents/internal/tools"
	"fmt"
	"testing"
)

// Implementing a fake model.

type MockModel struct {
	messages []model.AssistantMessage
	current  int
}

// Prompt implements [model.Model].
func (m *MockModel) Prompt(ctx context.Context, input []model.Message, tools model.Tools) (model.AssistantMessage, error) {
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

	registry := tools.NewRegistry()

	agent := NewAgent(&mockModel, uint8(10), *registry, "")

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
				`Looking up temperature in the bedroom`,
				[]model.ToolCall{
					model.NewToolCall(
						"1",
						"room_temperature",
						map[string]any{
							"room": "bedroom",
						},
					)},
			),
			model.NewAssistantMessage(
				`Looking up temperature in the living room`,
				[]model.ToolCall{
					model.NewToolCall(
						"2",
						"room_temperature",
						map[string]any{
							"room": "living_room",
						},
					)},
			),
		},
	}

	registry := tools.NewRegistry()
	temperature := tools.NewRoomTemperatureTool()
	registry.Register(temperature)
	agent := NewAgent(&mockModel, 2, *registry, "")
	r, err := agent.Run(t.Context(), "Check temperature in my bedroom and living room")

	if err == nil {
		t.Fatal("Expected run to return err, instead err is nil")
	}

	if r != `Looking up temperature in the living room` {
		t.Fatalf("Expected string: %s, instead it returned empty string", `Looking up temperature in the living room`)
	}

	if mockModel.current != 2 {
		t.Fatalf("Expected 2 calls, got %d", mockModel.current)
	}
}
