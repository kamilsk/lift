package paas

// A Redis contains configuration for a database.
type Redis struct {
	Enabled  *bool  `toml:"enabled"`
	Version  string `toml:"version"`
	Size     string `toml:"size"`
	Type     string `toml:"type,omitempty"`
	Replicas uint   `toml:"replicas,omitempty"`
}

// Merge combines two database configurations.
func (dst *Redis) Merge(src *Redis) {
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

	if src.Type != "" {
		dst.Type = src.Type
	}
	if src.Replicas != 0 {
		dst.Replicas = src.Replicas
	}
}

// A ShardedRedis contains configuration for a database.
type ShardedRedis struct {
	Enabled     *bool  `toml:"enabled"`
	Version     string `toml:"version"`
	Size        string `toml:"size"`
	Shards      Shards `toml:"shards"`
	SelfSharded *bool  `toml:"self-sharded,omitempty"`
}

// Merge combines two database configurations.
func (dst *ShardedRedis) Merge(src *ShardedRedis) {
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

	dst.Shards.Merge(src.Shards)
	if src.SelfSharded != nil {
		dst.SelfSharded = src.SelfSharded
	}
}
