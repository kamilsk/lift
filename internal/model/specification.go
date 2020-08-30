package model

type Specification struct {
	Name         string               `toml:"name,omitempty"`
	Description  string               `toml:"kind,omitempty"`
	Kind         string               `toml:"description,omitempty"`
	Host         string               `toml:"host,omitempty"`
	Replicas     uint                 `toml:"replicas,omitempty"`
	Engine       *Engine              `toml:"engine,omitempty"`
	Logger       *Logger              `toml:"logger,omitempty"`
	Balancer     *Balancer            `toml:"balancing,omitempty"`
	PostgreSQL   *PostgreSQL          `toml:"postgresql,omitempty"`
	Redis        *Redis               `toml:"redis,omitempty"`
	RedisSharded *ShardedRedis        `toml:"redis-sharded,omitempty"`
	SFTP         *SFTP                `toml:"sftp,omitempty"`
	Crons        Crons                `toml:"crons,omitepmty"`
	Dependencies Dependencies         `toml:"dependencies,omitempty"`
	Executable   Executable           `toml:"executable,omitempty"`
	Proxies      Proxies              `toml:"proxy,omitempty"`
	Queues       Queues               `toml:"queues,omitempty"`
	Sphinxes     Sphinxes             `toml:"sphinx,omitempty"`
	Workers      Workers              `toml:"workers,omitempty"`
	EnvVars      EnvironmentVariables `toml:"env_vars,omitempty"`
}

func (spec *Specification) Merge(src *Specification) {
	if spec == nil || src == nil {
		return
	}

	if src.Name != "" {
		spec.Name = src.Name
	}
	if src.Description != "" {
		spec.Description = src.Description
	}
	if src.Kind != "" {
		spec.Kind = src.Kind
	}
	if src.Host != "" {
		spec.Host = src.Host
	}
	if src.Replicas > 0 {
		spec.Replicas = src.Replicas
	}

	if src.Engine != nil && spec.Engine == nil {
		spec.Engine = new(Engine)
	}
	spec.Engine.Merge(src.Engine)

	if src.Logger != nil && spec.Logger == nil {
		spec.Logger = new(Logger)
	}
	spec.Logger.Merge(src.Logger)

	if src.Balancer != nil && spec.Balancer == nil {
		spec.Balancer = new(Balancer)
	}
	spec.Balancer.Merge(src.Balancer)

	if src.PostgreSQL != nil && spec.PostgreSQL == nil {
		spec.PostgreSQL = new(PostgreSQL)
	}
	spec.PostgreSQL.Merge(src.PostgreSQL)

	if src.Redis != nil && spec.Redis == nil {
		spec.Redis = new(Redis)
	}
	spec.Redis.Merge(src.Redis)

	if src.RedisSharded != nil && spec.RedisSharded == nil {
		spec.RedisSharded = new(ShardedRedis)
	}
	spec.RedisSharded.Merge(src.RedisSharded)

	if src.SFTP != nil && spec.SFTP == nil {
		spec.SFTP = new(SFTP)
	}
	spec.SFTP.Merge(src.SFTP)

	spec.Crons.Merge(src.Crons)
	spec.Dependencies.Merge(src.Dependencies)
	spec.Executable.Merge(src.Executable)
	spec.Proxies.Merge(src.Proxies)
	spec.Queues.Merge(src.Queues)
	spec.Sphinxes.Merge(src.Sphinxes)
	spec.Workers.Merge(src.Workers)
	spec.EnvVars.Merge(src.EnvVars)
}
