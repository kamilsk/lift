package model

import "sort"

// A DataBus contains configuration for data bus section of a PostgreSQL configuration.
type DataBus struct {
	BatchSize uint     `toml:"batch_size,omitempty"`
	Schemas   []string `toml:"schemas,omitempty"`
}

// Merge combines two data bus configurations.
func (databus *DataBus) Merge(src *DataBus) {
	if databus == nil || src == nil {
		return
	}

	if src.BatchSize != 0 {
		databus.BatchSize = src.BatchSize
	}
	if len(src.Schemas) > 0 {
		databus.Schemas = append(databus.Schemas, src.Schemas...)
	}

	if len(databus.Schemas) > 0 && !sort.StringsAreSorted(databus.Schemas) {
		sort.Strings(databus.Schemas)

		var current string
		unique := databus.Schemas[:0]
		for _, schema := range databus.Schemas {
			if current != schema {
				current = schema
				unique = append(unique, current)
			}
		}
		databus.Schemas = unique
	}
}
