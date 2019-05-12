package config

import (
	"io"
	"os"

	"github.com/kamilsk/platform/pkg/safe"
	"github.com/pelletier/go-toml"
)

// App contains service configuration.
type App struct {
	Name         string            `toml:"name"`
	Desc         string            `toml:"description"`
	Kind         string            `toml:"kind"`
	Unit         string            `toml:"unit"`
	Engine       Engine            `toml:"engine"`
	Environment  map[string]string `toml:"envs.local.env_vars"`
	Dependencies []Dependency      `toml:"dependencies"`
}

// Engine describes section related to a service engine.
type Engine struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Size    string `toml:"size"`
}

// Dependency describes section related to a service dependencies.
type Dependency struct {
	Name string `toml:"name"`
}

// Decode reads configuration from reader and decodes it into struct.
func Decode(r io.Reader) (App, error) {
	var cnf App
	err := toml.NewDecoder(r).Decode(&cnf)
	return cnf, err
}

// FromFile reads configuration from file and decodes it into struct.
func FromFile(file string) (App, error) {
	f, err := os.Open(file)
	if err != nil {
		return App{}, err
	}
	defer safe.Close(f)
	return Decode(f)
}
