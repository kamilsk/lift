package model

import "sort"

type Dependency struct {
	Name         string `toml:"name,omitempty"`
	Mock         bool   `toml:"mock,omitempty"`
	MockReplicas uint   `toml:"mock-replicas,omitempty"`
}

type Dependencies []Dependency

func (deps Dependencies) Len() int           { return len(deps) }
func (deps Dependencies) Less(i, j int) bool { return deps[i].Name < deps[j].Name }
func (deps Dependencies) Swap(i, j int)      { deps[i], deps[j] = deps[j], deps[i] }

func (deps *Dependencies) Merge(src Dependencies) {
	if deps == nil || len(src) == 0 {
		return
	}

	copied := *deps
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*deps = copied[:shift+1]
}
