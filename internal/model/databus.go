package model

import "go.octolab.org/strings"

// A DataBus contains configuration for data bus section of a PostgreSQL configuration.
type DataBus struct {
	BatchSize uint     `toml:"batch_size,omitempty"`
	Schemas   []string `toml:"schemas,omitempty"`
}

// Merge combines two data bus configurations.
func (dst *DataBus) Merge(src *DataBus) {
	if dst == nil || src == nil {
		return
	}

	if src.BatchSize != 0 {
		dst.BatchSize = src.BatchSize
	}
	if len(src.Schemas) > 0 {
		dst.Schemas = strings.Unique(append(dst.Schemas, src.Schemas...))
	}
}
