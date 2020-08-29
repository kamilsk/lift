package model

type PostgreSQL struct {
	Version  string `toml:"version,omitempty"`
	Size     string `toml:"size,omitempty"`
	Enabled  *bool  `toml:"enabled,omitempty"`
	OwnName  *bool  `toml:"use_own_maintenance_table_name,omitempty"`
	Fixtures *bool  `toml:"fixtures_enabled,omitempty"`
}

func (postgres *PostgreSQL) Merge(src *PostgreSQL) {
	if postgres == nil || src == nil {
		return
	}

	if src.Version != "" {
		postgres.Version = src.Version
	}
	if src.Size != "" {
		postgres.Size = src.Size
	}
	if src.Enabled != nil {
		postgres.Enabled = src.Enabled
	}
	if src.OwnName != nil {
		postgres.OwnName = src.OwnName
	}
	if src.Fixtures != nil {
		postgres.Fixtures = src.Fixtures
	}
}
