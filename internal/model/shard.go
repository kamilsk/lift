package model

import "go.octolab.org/strings"

// A Shard contains configuration for an abstract database shard.
type Shard struct {
	Master string   `toml:"master"`
	Slaves []string `toml:"slaves"`
}

// Shards is a list of Shard.
type Shards []Shard

// Len, Less, Swap implements the sort.Interface.
func (dst Shards) Len() int           { return len(dst) }
func (dst Shards) Less(i, j int) bool { return dst[i].Master < dst[j].Master }
func (dst Shards) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two shards configurations.
func (dst *Shards) Merge(src Shards) {
	if dst == nil || len(src) == 0 {
		return
	}

	copied := *dst
	copied = append(copied, src...)

	registry := map[string]int{}
	for i := len(copied); i > 0; i-- {
		registry[copied[i-1].Master] = i - 1
	}
	unique := copied[:0]
	for i, shard := range copied {
		origin := registry[shard.Master]
		if i == origin {
			unique = append(unique, shard)
			continue
		}
		unique[origin].Slaves = strings.Unique(append(unique[origin].Slaves, shard.Slaves...))
	}

	*dst = unique
}
