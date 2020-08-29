package model

type Redis struct {
	Version  string `toml:"version,omitempty"`
	Size     string `toml:"size,omitempty"`
	Type     string `toml:"type,omitempty"`
	Replicas int    `toml:"replicas,omitempty"`
	Enabled  *bool  `toml:"enabled,omitempty"`
}

func (redis *Redis) Merge(src *Redis) {
	if redis == nil || src == nil {
		return
	}

	if src.Version != "" {
		redis.Version = src.Version
	}
	if src.Size != "" {
		redis.Size = src.Size
	}
	if src.Type != "" {
		redis.Type = src.Type
	}
	if src.Replicas != 0 {
		redis.Replicas = src.Replicas
	}
	if src.Enabled != nil {
		redis.Enabled = src.Enabled
	}
}

type ShardedRedis struct {
	Version     string `toml:"version"`
	Size        string `toml:"size"`
	Shards      Shards `toml:"shards"`
	Enabled     bool   `toml:"enabled"`
	SelfSharded *bool  `toml:"self-sharded,omitempty"`
}

func (redis *ShardedRedis) Merge(src *ShardedRedis) {
	if redis == nil || src == nil {
		return
	}

	if src.Version != "" {
		redis.Version = src.Version
	}
	if src.Size != "" {
		redis.Size = src.Size
	}
	redis.Enabled = src.Enabled
	if src.SelfSharded != nil {
		redis.SelfSharded = src.SelfSharded
	}

	redis.Shards.Merge(src.Shards)
}
