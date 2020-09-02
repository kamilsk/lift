package model

// A Engine contains configuration for a service engine.
type Engine struct {
	Name      string     `toml:"name"`
	Version   string     `toml:"version"`
	Size      string     `toml:"size"`
	Resources *Resources `toml:"resources,omitempty"`
}

// Merge combines two service engine configurations.
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
