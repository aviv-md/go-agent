package tools

import (
	"reflect"
	"testing"
)

func TestToolCreationAndRun(t *testing.T) {

	tool := NewRoomTemperatureTool()
	schema := map[string]any{
		"type": "object",
		"properties": map[string]any{
			"room": map[string]any{
				"type":        "string",
				"description": "The room to get the temperature of",
			},
		},
		"required": []string{"room"},
	}
	res, err := tool.Run(map[string]any{"room": "kitchen"})
	if err != nil {
		t.Fatalf("Error running tool: %v", err)
	}

	if res != "The temperature in kitchen is 72 degrees" {
		t.Errorf("Expected 'The temperature in kitchen is 72 degrees', got '%s'", res)
	}

	eq := reflect.DeepEqual(schema, tool.Schema())
	if !eq {
		t.Fatalf("Expected schema to be equal, got %v", tool.schema)
	}
}
