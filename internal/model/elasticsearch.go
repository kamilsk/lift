package model

// A ElasticSearch contains configuration for a database.
type ElasticSearch struct {
	Enabled *bool  `toml:"enabled"`
	Version string `toml:"version"`
	Size    string `toml:"size"`
}

// Merge combines two database configurations.
func (dst *ElasticSearch) Merge(src *ElasticSearch) {
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
