package config

import (
	"io"
	"os"
	"strings"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/pelletier/go-toml"

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

	var cnf extended
	err := toml.NewDecoder(r).Decode(&cnf)
	if err != nil {
		return Service{}, err
	}
	cnf.Engine.WorkDir = scope.WorkingDir
	for name, storage := range map[string]*storage{
		storageMongo:    &cnf.MongoDB,
		storagePostgres: &cnf.PostgreSQL,
		storageRabbit:   &cnf.RabbitMQ,
		storageRedis:    &cnf.Redis,
		storageSphinx:   &cnf.Sphinx,
	} {
		if storage.Enabled {
			cnf.Dependencies = append(cnf.Dependencies, defaults.FindByName(name))
		}
	}
	env := Environment{}
	for _, dep := range cnf.Dependencies {
		env.Merge(dep.vars)
	}
	env.Merge(EngineSpecific(cnf.Engine))
	env.Merge(cnf.Environment)
	env.Merge(cnf.Local)
	return Service{
		Name:         cnf.Name,
		Engine:       cnf.Engine,
		Environment:  env,
		Dependencies: cnf.Dependencies,
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
	defer safe.Close(f)
	return Decode(scope, f)
}
