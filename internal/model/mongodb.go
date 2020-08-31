package model

type MongoDB struct {
	Version string `toml:"version,omitempty"`
	Size    string `toml:"size,omitempty"`
	Enabled *bool  `toml:"enabled,omitempty"`
}

func (mongodb *MongoDB) Merge(src *MongoDB) {
	if mongodb == nil || src == nil {
		return
	}

	if src.Version != "" {
		mongodb.Version = src.Version
	}
	if src.Size != "" {
		mongodb.Size = src.Size
	}
	if src.Enabled != nil {
		mongodb.Enabled = src.Enabled
	}
}
