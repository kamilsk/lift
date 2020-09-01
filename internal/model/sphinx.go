package model

import "sort"

// A Sphinx contains configuration for a database.
type Sphinx struct {
	Enabled *bool  `toml:"enabled"`
	Hosts   Hosts  `toml:"hosts"`
	Name    string `toml:"name,omitepmty"`
	Haproxy string `toml:"haproxy_tag,omitempty"`
}

// Sphinxes is a list of Sphinx.
type Sphinxes []Sphinx

// Len, Less, Swap implements the sort.Interface.
func (dst Sphinxes) Len() int           { return len(dst) }
func (dst Sphinxes) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Sphinxes) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two sphinxes configurations.
func (dst *Sphinxes) Merge(src Sphinxes) {
	if dst == nil || len(src) == 0 {
		return
	}

	copied := *dst
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
	*dst = copied[:shift+1]
}
