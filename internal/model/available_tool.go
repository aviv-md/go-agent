package model

type AvailableTool struct {
	name        string
	description string
	parameters  map[string]any
}

func (a AvailableTool) Name() string {
	return a.name
}

func (a AvailableTool) Description() string {
	return a.description
}

func (a AvailableTool) Parameters() map[string]any {
	return a.parameters
}

func NewAvailableTool(name string, description string, parameters map[string]any) AvailableTool {
	return AvailableTool{
		name:        name,
		description: description,
		parameters:  parameters,
	}
}
