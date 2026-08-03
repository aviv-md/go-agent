package tools

import (
	"testing"
)

func TestRegistryCreationAndRegistration(t *testing.T) {
	trt := NewRoomTemperatureTool()

	reg := NewRegistry()

	reg.Register(trt)

	res, err := reg.Execute(trt.Name(), map[string]any{"room": "kitchen"})
	if err != nil {
		t.Fatalf("Error running tool: %v", err)
	}

	if res != "The temperature in kitchen is 72 degrees" {
		t.Errorf("Expected 'The temperature in kitchen is 72 degrees', got '%s'", res)
	}
}

func TestUnknownTool(t *testing.T) {
	reg := NewRegistry()

	_, err := reg.Execute("unknown_tool", map[string]any{"room": "kitchen"})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestListTools(t *testing.T) {
	reg := NewRegistry()

	tFn := func(args map[string]any) (string, error) {
		return "test", nil
	}

	tName := "test_tool"
	tDesc := "Test tool"
	tParameters := map[string]any{
		"type": "object",
	}

	trt := NewTool(tName, tDesc, tParameters, tFn)

	reg.Register(trt)

	tools := reg.List()

	if len(tools) != 1 {
		t.Fatalf("Expected 1 tool, got %d", len(tools))
	}

	if tools[0].Name() != tName {
		t.Errorf("Expected tool name to be %s", tName)
	}
}
