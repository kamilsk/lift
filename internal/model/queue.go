package model

import "go.octolab.org/strings"

// A Queue contains configuration for a message queue.
type Queue struct {
	Name    string   `toml:"schema"`
	DLQ     []string `toml:"dlq,omitempty"`
	Aliases []string `toml:"aliases,omitempty"`
}

// Merge combines two message queue configurations.
func (dst *Queue) Merge(src Queue) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if len(src.DLQ) > 0 {
		dst.DLQ = strings.Unique(append(dst.DLQ, src.DLQ...))
	}
	if len(src.Aliases) > 0 {
		dst.Aliases = strings.Unique(append(dst.Aliases, src.Aliases...))
	}
}

// Queues is a list of Queue.
type Queues []Queue

// Len, Less, Swap implements the sort.Interface.
func (dst Queues) Len() int           { return len(dst) }
func (dst Queues) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Queues) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of message queue configurations.
func (dst *Queues) Merge(src Queues) {
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
	for i, queue := range copied {
		origin := registry[queue.Name]
		if i == origin {
			unique = append(unique, queue)
			continue
		}
		unique[origin].Merge(queue)
	}

	*dst = unique
}
