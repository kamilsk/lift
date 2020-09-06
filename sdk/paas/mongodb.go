package paas

// A MongoDB contains configuration for a database.
type MongoDB struct {
	Enabled *bool  `toml:"enabled"`
	Version string `toml:"version"`
	Size    string `toml:"size"`
}

// Merge combines two database configurations.
func (dst *MongoDB) Merge(src *MongoDB) {
	if dst == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Version != "" {
		dst.Version = src.Version
	}
	if src.Size != "" {
		dst.Size = src.Size
	}
}
