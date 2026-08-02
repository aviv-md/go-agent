package tools

import "fmt"

func NewRoomTemperatureTool() Tool {
	fn := func(args map[string]any) (string, error) {
		r, ok := args["room"].(string)
		if !ok {
			return "", fmt.Errorf("room must be a string")
		}
		return fmt.Sprintf("The temperature in %s is 72 degrees", r), nil
	}

	name := "room_temperature"
	desc := "Get the temperature of a room"
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

	return NewTool(
		name,
		desc,
		schema,
		fn,
	)
}
