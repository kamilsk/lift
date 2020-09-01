package model

// A PostgreSQL contains configuration for a database.
type PostgreSQL struct {
	Enabled  *bool    `toml:"enabled"`
	Version  string   `toml:"version"`
	Size     string   `toml:"size"`
	Fixtures *bool    `toml:"fixtures_enabled,omitempty"`
	OwnName  *bool    `toml:"use_own_maintenance_table_name,omitempty"`
	DataBus  *DataBus `toml:"data_bus,omitempty"`
}

// Merge combines two database configurations.
func (dst *PostgreSQL) Merge(src *PostgreSQL) {
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

	if src.Fixtures != nil {
		dst.Fixtures = src.Fixtures
	}
	if src.OwnName != nil {
		dst.OwnName = src.OwnName
	}

	if src.DataBus != nil && dst.DataBus == nil {
		dst.DataBus = new(DataBus)
	}
	dst.DataBus.Merge(src.DataBus)
}
