package model

import "sort"

type Queue struct {
	Name    string   `toml:"schema,omitempty"`
	DLQ     []string `toml:"dlq,omitempty"`
	Aliases []string `toml:"aliases,omitempty"`
}

type Queues []Queue

func (queues Queues) Len() int           { return len(queues) }
func (queues Queues) Less(i, j int) bool { return queues[i].Name < queues[j].Name }
func (queues Queues) Swap(i, j int)      { queues[i], queues[j] = queues[j], queues[i] }

func (queues *Queues) Merge(src Queues) {
	if queues == nil || len(src) == 0 {
		return
	}

	copied := *queues
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
	*queues = copied[:shift+1]
}
