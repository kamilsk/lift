package model

import "sort"

type Shard struct {
	Master string   `toml:"master"`
	Slaves []string `toml:"slaves"`
}

type Shards []Shard

func (shards Shards) Len() int           { return len(shards) }
func (shards Shards) Less(i, j int) bool { return shards[i].Master < shards[j].Master }
func (shards Shards) Swap(i, j int)      { shards[i], shards[j] = shards[j], shards[i] }

func (shards *Shards) Merge(src Shards) {
	if shards == nil || len(src) == 0 {
		return
	}

	copied := *shards
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Master == copied[i].Master {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*shards = copied[:shift+1]
}
