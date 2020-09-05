package model

import "go.octolab.org/strings"

// A Shard contains configuration for an abstract database shard.
type Shard struct {
	Primary string   `toml:"master"`
	Reserve []string `toml:"slaves"`
}

// Merge combines two shard configurations.
func (dst *Shard) Merge(src Shard) {
	if dst == nil || dst.Primary != src.Primary {
		return
	}

	if len(src.Reserve) > 0 {
		dst.Reserve = strings.Unique(append(dst.Reserve, src.Reserve...))
	}
}

// Shards is a list of Shard.
type Shards []Shard

// Len, Less, Swap implements the sort.Interface.
func (dst Shards) Len() int           { return len(dst) }
func (dst Shards) Less(i, j int) bool { return dst[i].Primary < dst[j].Primary }
func (dst Shards) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of shard configurations.
func (dst *Shards) Merge(src Shards) {
	if dst == nil || len(src) == 0 {
		return
	}

	copied := *dst
	copied = append(copied, src...)

	registry := map[string]int{}
	for i := len(copied); i > 0; i-- {
		registry[copied[i-1].Primary] = i - 1
	}
	unique := copied[:0]
	for i, shard := range copied {
		origin := registry[shard.Primary]
		if i == origin {
			unique = append(unique, shard)
			continue
		}
		unique[origin].Merge(shard)
	}

	*dst = unique
}
