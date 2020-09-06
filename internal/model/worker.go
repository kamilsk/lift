package model

import (
	"go.octolab.org/strings"
)

// A Worker contains configuration for unit of work.
type Worker struct {
	Name          string     `toml:"name"`
	Enabled       *bool      `toml:"enabled,omitempty"`
	Command       string     `toml:"command,omitempty"`
	Commands      []string   `toml:"commands,omitempty"`
	Replicas      uint       `toml:"replicas"`
	LivenessProbe string     `toml:"liveness-probe-command"`
	Size          string     `toml:"size"`
	Resources     *Resources `toml:"resources,omitempty"`
}

// Merge combines two unit of work configurations.
func (dst *Worker) Merge(src Worker) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Command != "" {
		dst.Command = src.Command
	}
	if len(src.Commands) > 0 {
		dst.Commands = strings.Unique(append(dst.Commands, src.Commands...))
	}
	if src.Replicas != 0 {
		dst.Replicas = src.Replicas
	}
	if src.LivenessProbe != "" {
		dst.LivenessProbe = src.LivenessProbe
	}
	if src.Size != "" {
		dst.Size = src.Size
	}

	if src.Resources != nil && dst.Resources == nil {
		dst.Resources = new(Resources)
	}
	dst.Resources.Merge(src.Resources)
}

// Workers is a list of Worker.
type Workers []Worker

// Len, Less, Swap implements the sort.Interface.
func (dst Workers) Len() int           { return len(dst) }
func (dst Workers) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Workers) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of unit of work configurations.
func (dst *Workers) Merge(src Workers) {
	if dst == nil || len(src) == 0 {
		return
	}

	copied := *dst
	copied = append(copied, src...)

	registry := map[string]int{}
	for i := len(copied); i > 0; i-- {
		registry[copied[i-1].Name] = i - 1
	}
	unique := copied[:0]
	for i, worker := range copied {
		origin := registry[worker.Name]
		if i == origin {
			unique = append(unique, worker)
			continue
		}
		unique[origin].Merge(worker)
	}

	*dst = unique
}
