package model

type ElasticSearch struct {
	Enabled *bool  `toml:"enabled,omitempty"`
	Version string `toml:"version,omitempty"`
	Size    string `toml:"size,omitempty"`
}

func (elastic *ElasticSearch) Merge(src *ElasticSearch) {
	if elastic == nil || src == nil {
		return
	}

	if src.Enabled != nil {
		elastic.Enabled = src.Enabled
	}
	if src.Version != "" {
		elastic.Version = src.Version
	}
	if src.Size != "" {
		elastic.Size = src.Size
	}
}
