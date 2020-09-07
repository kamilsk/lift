package paas

// A Engine contains configuration for a service engine.
type Engine struct {
	Name      string     `toml:"name,omitempty"`
	Version   string     `toml:"version,omitempty"`
	Size      string     `toml:"size,omitempty"`
	Resources *Resources `toml:"resources,omitempty"`
}

// Merge combines two service engine configurations.
func (dst *Engine) Merge(src *Engine) {
	if dst == nil || src == nil {
		return
	}

	if src.Name != "" {
		dst.Name = src.Name
	}
	if src.Version != "" {
		dst.Version = src.Version
	}
	if src.Size != "" {
		dst.Size = src.Size
	}

	if src.Resources != nil && dst.Resources == nil {
		dst.Resources = new(Resources)
	}
	dst.Resources.Merge(src.Resources)
}
