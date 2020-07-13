package model

type Application struct {
	Specification `toml:",omitempty,squash"`
	Envs          map[string]Specification `toml:"envs,omitempty"`
}

type Balancing struct {
	CookieAffinity string `toml:"cookie_affinity,omitempty"`
}

type Cron struct {
	Name     string `toml:"name,omitempty"`
	Enabled  bool   `toml:"enabled,omitempty"`
	Schedule string `toml:"schedule,omitempty"`
	Command  string `toml:"command,omitempty"`
}

type Crons []Cron

func (crons Crons) Len() int           { return len(crons) }
func (crons Crons) Less(i, j int) bool { return crons[i].Name < crons[j].Name }
func (crons Crons) Swap(i, j int)      { crons[i], crons[j] = crons[j], crons[i] }

type Dependency struct {
	Name     string `toml:"name,omitempty"`
	Mock     bool   `toml:"mock,omitempty"`
	Replicas uint   `toml:"mock-replicas,omitempty"`
}

type Dependencies []Dependency

func (deps Dependencies) Len() int           { return len(deps) }
func (deps Dependencies) Less(i, j int) bool { return deps[i].Name < deps[j].Name }
func (deps Dependencies) Swap(i, j int)      { deps[i], deps[j] = deps[j], deps[i] }

type Engine struct {
	Name      string   `toml:"name,omitempty"`
	Version   string   `toml:"version,omitempty"`
	Size      string   `toml:"size,omitempty"`
	Resources Resource `toml:"resources,omitempty"`
}

type EnvironmentVariables map[string]string

type Exec struct {
	Name          string `toml:"name,omitempty"`
	Replicas      uint   `toml:"replicas,omitempty"`
	Command       string `toml:"command,omitempty"`
	Port          string `toml:"service-port,omitempty"`
	Size          string `toml:"size,omitempty"`
	RedinessProbe string `toml:"readiness-probe-command,omitempty"`
	LivenessProbe string `toml:"liveness-probe-command,omitempty"`
}

type Executable []Exec

func (exec Executable) Len() int           { return len(exec) }
func (exec Executable) Less(i, j int) bool { return exec[i].Name < exec[j].Name }
func (exec Executable) Swap(i, j int)      { exec[i], exec[j] = exec[j], exec[i] }

type Host struct {
	Name        string `toml:"host,omitempty"`
	Connections uint   `toml:"connections,omitempty"`
	Weight      uint   `toml:"weight,omitempty"`
	Backup      bool   `toml:"backup,omitempty"`
}

type Hosts []Host

func (hosts Hosts) Len() int           { return len(hosts) }
func (hosts Hosts) Less(i, j int) bool { return hosts[i].Name < hosts[j].Name }
func (hosts Hosts) Swap(i, j int)      { hosts[i], hosts[j] = hosts[j], hosts[i] }

type Logger struct {
	Level string `toml:"level,omitempty"`
}

type Proxy struct {
	Name    string `toml:"name,omitempty"`
	Enabled bool   `toml:"enabled,omitempty"`
	Hosts   Hosts  `toml:"hosts,omitempty"`
}

type Proxies []Proxy

func (proxies Proxies) Len() int           { return len(proxies) }
func (proxies Proxies) Less(i, j int) bool { return proxies[i].Name < proxies[j].Name }
func (proxies Proxies) Swap(i, j int)      { proxies[i], proxies[j] = proxies[j], proxies[i] }

type Queue struct {
	Schema  string   `toml:"schema,omitempty"`
	DLQ     []string `toml:"dlq,omitempty"`
	Aliases []string `toml:"aliases,omitempty"`
}

type Queues []Queue

func (queues Queues) Len() int           { return len(queues) }
func (queues Queues) Less(i, j int) bool { return queues[i].Schema < queues[j].Schema }
func (queues Queues) Swap(i, j int)      { queues[i], queues[j] = queues[j], queues[i] }

type Resource struct {
	CPU    int `toml:"cpu,omitempty"`
	Memory int `toml:"memory,omitempty"`
}

type Resources struct {
	Requests Resource `toml:"requests,omitempty"`
	Limits   Resource `toml:"limits,omitempty"`
}

type Specification struct {
	Name         string               `toml:"name,omitempty"`
	Description  string               `toml:"kind,omitempty"`
	Kind         string               `toml:"description,omitempty"`
	Host         string               `toml:"host,omitempty"`
	Replicas     uint                 `toml:"replicas,omitempty"`
	Engine       Engine               `toml:"engine,omitempty"`
	Logger       Logger               `toml:"logger,omitempty"`
	Balancing    Balancing            `toml:"balancing,omitempty"`
	SFTP         SFTP                 `toml:"sftp,omitempty"`
	Crons        Crons                `toml:"crons,omitepmty"`
	Dependencies Dependencies         `toml:"dependencies,omitempty"`
	Executable   Executable           `toml:"executable,omitempty"`
	Proxies      Proxies              `toml:"proxy,omitempty"`
	Queues       Queues               `toml:"queues,omitempty"`
	Workers      Workers              `toml:"workers,omitempty"`
	EnvVars      EnvironmentVariables `toml:"env_vars,omitempty"`
}

type SFTP struct {
	Size    string `toml:"size,omitempty"`
	Enabled bool   `toml:"enabled,omitempty"`
}

type Worker struct {
	Name          string    `toml:"name,omitempty"`
	Replicas      uint      `toml:"replicas,omitempty"`
	Command       string    `toml:"command,omitempty"`
	Commands      []string  `toml:"commands,omitempty"`
	Size          string    `toml:"size,omitempty"`
	Enabled       bool      `toml:"enabled,omitempty"`
	LivenessProbe string    `toml:"liveness-probe-command,omitempty"`
	Resources     Resources `toml:"resources,omitempty"`
}

type Workers []Worker

func (workers Workers) Len() int           { return len(workers) }
func (workers Workers) Less(i, j int) bool { return workers[i].Name < workers[j].Name }
func (workers Workers) Swap(i, j int)      { workers[i], workers[j] = workers[j], workers[i] }
