package model

type Application struct {
	Specification `toml:",omitempty,squash"`
	Envs          map[string]*Specification `toml:"envs,omitempty"`
}

type EnvironmentVariables map[string]string

type Host struct {
	Name        string `toml:"host,omitempty"`
	AgentPort   uint   `toml:"agent_port,omitempty"`
	Connections uint   `toml:"connections,omitempty"`
	MaxConns    uint   `toml:"maxconn,omitempty"`
	Weight      uint   `toml:"weight,omitempty"`
	Backup      bool   `toml:"backup,omitempty"`
}

type Hosts []Host

func (hosts Hosts) Len() int           { return len(hosts) }
func (hosts Hosts) Less(i, j int) bool { return hosts[i].Name < hosts[j].Name }
func (hosts Hosts) Swap(i, j int)      { hosts[i], hosts[j] = hosts[j], hosts[i] }

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

type Worker struct {
	Name          string     `toml:"name,omitempty"`
	Enabled       *bool      `toml:"enabled,omitempty"`
	Replicas      uint       `toml:"replicas,omitempty"`
	Command       string     `toml:"command,omitempty"`
	Commands      []string   `toml:"commands,omitempty"`
	Size          string     `toml:"size,omitempty"`
	LivenessProbe string     `toml:"liveness-probe-command,omitempty"`
	Resources     *Resources `toml:"resources,omitempty"`
}

type Workers []Worker

func (workers Workers) Len() int           { return len(workers) }
func (workers Workers) Less(i, j int) bool { return workers[i].Name < workers[j].Name }
func (workers Workers) Swap(i, j int)      { workers[i], workers[j] = workers[j], workers[i] }
