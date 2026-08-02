package tools

import "fmt"

type Registry struct {
	tools map[string]Tool
}

func (r *Registry) Register(tool Tool) {
	r.tools[tool.Name()] = tool
}

func (r *Registry) Execute(name string, args map[string]any) (string, error) {
	tool, ok := r.tools[name]
	if !ok {
		return "", fmt.Errorf("tool %s not found", name)
	}
	return tool.Run(args)
}

func (r *Registry) List() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}
