package model

type Dependency struct {
	Name     string `toml:"name"`
	Mock     bool   `toml:"mock,omitempty"`
	Replicas uint   `toml:"mock-replicas,omitempty"`
}

type Dependencies []Dependency

func (deps Dependencies) Len() int           { return len(deps) }
func (deps Dependencies) Less(i, j int) bool { return deps[i].Name < deps[j].Name }
func (deps Dependencies) Swap(i, j int)      { deps[i], deps[j] = deps[j], deps[i] }
