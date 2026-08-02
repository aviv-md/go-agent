package tools

type (
	Tool struct {
		name        string
		description string
		parameters  map[string]any
		fn          func(args map[string]any) (string, error)
	}
)

func (t Tool) Description() string {
	return t.description
}

func (t Tool) Parameters() map[string]any {
	return t.parameters
}

func (t Tool) Name() string {
	return t.name
}

func (t Tool) Run(args map[string]any) (string, error) {
	return t.fn(args)
}

func NewTool(
	name string,
	description string,
	parameters map[string]any,
	fn func(args map[string]any) (string, error),
) Tool {
	return Tool{
		name:        name,
		description: description,
		parameters:  parameters,
		fn:          fn,
	}
}
