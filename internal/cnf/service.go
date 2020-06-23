package cnf

import (
	"io"
	"os"
	"strings"

	"github.com/pelletier/go-toml"
	"go.octolab.org/safe"
	"go.octolab.org/unsafe"

	"github.com/kamilsk/lift/internal"
)

// Service contains service configuration.
type Service struct {
	Name         string            `toml:"name"`
	Engine       Engine            `toml:"engine"`
	Environment  Environment       `toml:"env_vars"`
	Dependencies Dependencies      `toml:"dependencies"`
	PortMapping  map[uint16]uint16 `toml:"-"`
}

// Engine describes section related to a service engine.
type Engine struct {
	Name    string `toml:"name"`
	Size    string `toml:"size"`
	Version string `toml:"version"`
	WorkDir string `toml:"-"`
}

// Environment holds available environment variables.
type Environment map[string]string

// Merge merges new variables into present.
func (env Environment) Merge(extra Environment) {
	for k, v := range extra {
		env[k] = v
	}
}

// Dependency describes section related to a service dependencies.
type Dependency struct {
	Name    string   `toml:"name"`
	Forward []string `toml:"-"`

	vars map[string]string
}

// Dependencies holds service dependencies.
type Dependencies []Dependency

// FindByName finds dependency by its name or returns default empty value.
func (deps Dependencies) FindByName(name string) Dependency {
	for _, dep := range deps {
		if strings.EqualFold(name, dep.Name) {
			return dep
		}
	}
	return Dependency{}
}

type storage struct {
	Enabled bool   `toml:"enabled"`
	Size    string `toml:"size"`
	Version string `toml:"version"`
}

// Decode reads configuration from reader and decodes it into the struct.
func Decode(scope internal.Scope, r io.Reader) (Service, error) {
	type extended struct {
		Service

		Desc string `toml:"description"`
		Host string `toml:"host"`
		Kind string `toml:"kind"`

		MongoDB    storage `toml:"mongodb"`
		PostgreSQL storage `toml:"postgresql"`
		RabbitMQ   storage `toml:"rabbitmq"`
		Redis      storage `toml:"redis"`
		Sphinx     storage `toml:"sphinx"`

		Local Environment `toml:"envs.local.env_vars"`
	}

	var config extended
	err := toml.NewDecoder(r).Decode(&config)
	if err != nil {
		return Service{}, err
	}
	config.Engine.WorkDir = scope.WorkingDir
	for name, storage := range map[string]*storage{
		storageMongo:    &config.MongoDB,
		storagePostgres: &config.PostgreSQL,
		storageRabbit:   &config.RabbitMQ,
		storageRedis:    &config.Redis,
		storageSphinx:   &config.Sphinx,
	} {
		if storage.Enabled {
			config.Dependencies = append(config.Dependencies, defaults.FindByName(name))
		}
	}
	env := Environment{}
	for _, dep := range config.Dependencies {
		env.Merge(dep.vars)
	}
	env.Merge(EngineSpecific(config.Engine))
	env.Merge(config.Environment)
	env.Merge(config.Local)
	return Service{
		Name:         config.Name,
		Engine:       config.Engine,
		Environment:  env,
		Dependencies: config.Dependencies,
		PortMapping:  scope.PortMapping,
	}, nil
}

// FromScope reads configuration from file and decodes it into the struct.
func FromScope(scope internal.Scope, err error) (Service, error) {
	var service Service
	if err != nil {
		return service, err
	}
	f, err := os.Open(scope.ConfigPath)
	if err != nil {
		return service, err
	}
	defer safe.Close(f, unsafe.Ignore)
	return Decode(scope, f)
}
