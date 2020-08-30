package model

import "sort"

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

func (workers *Workers) Merge(src Workers) {
	if workers == nil || len(src) == 0 {
		return
	}

	copied := *workers
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
	*workers = copied[:shift+1]
}
