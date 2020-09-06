package model

// An Specification contains configuration for a service in a specific environment.
type Specification struct {
	Name        string `toml:"name,omitempty"`
	Description string `toml:"kind,omitempty"`
	Kind        string `toml:"description,omitempty"`
	Host        string `toml:"host,omitempty"`
	Replicas    uint   `toml:"replicas,omitempty"`

	Balancer *Balancer `toml:"balancing,omitempty"`
	Engine   *Engine   `toml:"engine,omitempty"`
	Logger   *Logger   `toml:"logger,omitempty"`
	SFTP     *SFTP     `toml:"sftp,omitempty"`

	Elastic      *ElasticSearch `toml:"elasticsearch,omitempty"`
	MongoDB      *MongoDB       `toml:"mongodb,omitempty"`
	PostgreSQL   *PostgreSQL    `toml:"postgresql,omitempty"`
	RabbitMQ     *RabbitMQ      `toml:"rabbitmq,omitempty"`
	Redis        *Redis         `toml:"redis,omitempty"`
	RedisSharded *ShardedRedis  `toml:"redis-sharded,omitempty"`
	Sphinxes     Sphinxes       `toml:"sphinx,omitempty"`

	Crons        Crons                `toml:"crons,omitepmty"`
	Dependencies Dependencies         `toml:"dependencies,omitempty"`
	EnvVars      EnvironmentVariables `toml:"env_vars,omitempty"`
	Executable   Executable           `toml:"executable,omitempty"`
	Proxies      Proxies              `toml:"proxy,omitempty"`
	Queues       Queues               `toml:"queues,omitempty"`
	Workers      Workers              `toml:"workers,omitempty"`
}

// Merge combines two service configurations in a specific environment.
func (spec *Specification) Merge(src *Specification) {
	if spec == nil || src == nil {
		return
	}

	// general

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
	if src.Replicas != 0 {
		spec.Replicas = src.Replicas
	}

	// misc

	if src.Balancer != nil && spec.Balancer == nil {
		spec.Balancer = new(Balancer)
	}
	spec.Balancer.Merge(src.Balancer)

	if src.Engine != nil && spec.Engine == nil {
		spec.Engine = new(Engine)
	}
	spec.Engine.Merge(src.Engine)

	if src.Logger != nil && spec.Logger == nil {
		spec.Logger = new(Logger)
	}
	spec.Logger.Merge(src.Logger)

	if src.SFTP != nil && spec.SFTP == nil {
		spec.SFTP = new(SFTP)
	}
	spec.SFTP.Merge(src.SFTP)

	// databases

	if src.Elastic != nil && spec.Elastic == nil {
		spec.Elastic = new(ElasticSearch)
	}
	spec.Elastic.Merge(src.Elastic)

	if src.MongoDB != nil && spec.MongoDB == nil {
		spec.MongoDB = new(MongoDB)
	}
	spec.MongoDB.Merge(src.MongoDB)

	if src.PostgreSQL != nil && spec.PostgreSQL == nil {
		spec.PostgreSQL = new(PostgreSQL)
	}
	spec.PostgreSQL.Merge(src.PostgreSQL)

	if src.RabbitMQ != nil && spec.RabbitMQ == nil {
		spec.RabbitMQ = new(RabbitMQ)
	}
	spec.RabbitMQ.Merge(src.RabbitMQ)

	if src.Redis != nil && spec.Redis == nil {
		spec.Redis = new(Redis)
	}
	spec.Redis.Merge(src.Redis)

	if src.RedisSharded != nil && spec.RedisSharded == nil {
		spec.RedisSharded = new(ShardedRedis)
	}
	spec.RedisSharded.Merge(src.RedisSharded)

	spec.Sphinxes.Merge(src.Sphinxes)

	// sets

	spec.Crons.Merge(src.Crons)
	spec.Dependencies.Merge(src.Dependencies)
	spec.EnvVars.Merge(src.EnvVars)
	spec.Executable.Merge(src.Executable)
	spec.Proxies.Merge(src.Proxies)
	spec.Queues.Merge(src.Queues)
	spec.Workers.Merge(src.Workers)
}
