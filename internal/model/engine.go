package model

type Engine struct {
	Name      string     `toml:"name,omitempty"`
	Version   string     `toml:"version,omitempty"`
	Size      string     `toml:"size,omitempty"`
	Resources *Resources `toml:"resources,omitempty"`
}

func (engine *Engine) Merge(src *Engine) {
	if engine == nil || src == nil {
		return
	}

	if src.Name != "" {
		engine.Name = src.Name
	}
	if src.Version != "" {
		engine.Version = src.Version
	}
	if src.Size != "" {
		engine.Size = src.Size
	}

	if src.Resources != nil && engine.Resources == nil {
		engine.Resources = new(Resources)
	}
	engine.Resources.Merge(src.Resources)
}
