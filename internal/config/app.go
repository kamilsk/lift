package config

import _ "github.com/pelletier/go-toml"

type App struct {
	Name         string       `toml:"name"`
	Desc         string       `toml:"description"`
	Kind         string       `toml:"kind"`
	Unit         string       `toml:"unit"`
	Engine       Engine       `toml:"engine"`
	Dependencies []Dependency `toml:"dependencies"`
}

type Engine struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Size    string `toml:"size"`
}

type Dependency struct {
	Name string `toml:"name"`
}
