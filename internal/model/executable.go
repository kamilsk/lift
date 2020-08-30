package model

import "sort"

type Exec struct {
	Name          string `toml:"name,omitempty"`
	Replicas      uint   `toml:"replicas,omitempty"`
	Command       string `toml:"command,omitempty"`
	Port          uint   `toml:"service-port,omitempty"`
	Size          string `toml:"size,omitempty"`
	RedinessProbe string `toml:"readiness-probe-command,omitempty"`
	LivenessProbe string `toml:"liveness-probe-command,omitempty"`
}

type Executable []Exec

func (exec Executable) Len() int           { return len(exec) }
func (exec Executable) Less(i, j int) bool { return exec[i].Name < exec[j].Name }
func (exec Executable) Swap(i, j int)      { exec[i], exec[j] = exec[j], exec[i] }

func (exec *Executable) Merge(src Executable) {
	if exec == nil || len(src) == 0 {
		return
	}

	copied := *exec
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
	*exec = copied[:shift+1]
}
