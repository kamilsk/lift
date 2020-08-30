package model

import "sort"

type Sphinx struct {
	Name    string `toml:"name,omitepmty"`
	Enabled *bool  `toml:"enabled,omitepmty"`
	Haproxy string `toml:"haproxy_tag,omitempty"`
	Hosts   Hosts  `toml:"hosts,omitempty"`
}

type Sphinxes []Sphinx

func (sphinxes Sphinxes) Len() int           { return len(sphinxes) }
func (sphinxes Sphinxes) Less(i, j int) bool { return sphinxes[i].Name < sphinxes[j].Name }
func (sphinxes Sphinxes) Swap(i, j int)      { sphinxes[i], sphinxes[j] = sphinxes[j], sphinxes[i] }

func (sphinxes *Sphinxes) Merge(src Sphinxes) {
	if sphinxes == nil || len(src) == 0 {
		return
	}

	copied := *sphinxes
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
	*sphinxes = copied[:shift+1]
}
