package config

import (
	"io"
	"os"
	"strings"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/pelletier/go-toml"
)

// Service contains service configuration.
type Service struct {
	Name         string       `toml:"name"`
	Desc         string       `toml:"description"`
	Host         string       `toml:"host"`
	Kind         string       `toml:"kind"`
	Unit         string       `toml:"unit"`
	Engine       Engine       `toml:"engine"`
	Environment  Environment  `toml:"env_vars"`
	Dependencies Dependencies `toml:"dependencies"`
}

// Engine describes section related to a service engine.
type Engine struct {
	Name    string `toml:"name"`
	Size    string `toml:"size"`
	Version string `toml:"version"`
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
	Name string `toml:"name"`

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

// Decode reads configuration from reader and decodes it into struct.
func Decode(r io.Reader) (Service, error) {
	type extended struct {
		Service
		MongoDB    storage     `toml:"mongodb"`
		PostgreSQL storage     `toml:"postgresql"`
		RabbitMQ   storage     `toml:"rabbitmq"`
		Redis      storage     `toml:"redis"`
		Sphinx     storage     `toml:"sphinx"`
		Local      Environment `toml:"envs.local.env_vars"`
	}

	var cnf extended
	err := toml.NewDecoder(r).Decode(&cnf)
	if err != nil {
		return Service{}, err
	}
	for name, storage := range map[string]*storage{
		mongoStorage:  &cnf.MongoDB,
		pgStorage:     &cnf.PostgreSQL,
		rabbitStorage: &cnf.RabbitMQ,
		redisStorage:  &cnf.Redis,
		sphinxStorage: &cnf.Sphinx,
	} {
		if storage.Enabled {
			cnf.Service.Dependencies = append(cnf.Service.Dependencies, Dependency{
				Name: name,
				vars: defaults.FindByName(name).vars,
			})
		}
	}
	env := make(Environment)
	for _, dep := range cnf.Dependencies {
		env.Merge(dep.vars)
	}
	env.Merge(cnf.Service.Environment)
	env.Merge(cnf.Local)
	cnf.Service.Environment = env
	return cnf.Service, nil
}

// FromFile reads configuration from file and decodes it into struct.
func FromFile(file string) (Service, error) {
	f, err := os.Open(file)
	if err != nil {
		return Service{}, err
	}
	defer safe.Close(f)
	return Decode(f)
}
