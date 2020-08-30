package model

import "sort"

type Proxy struct {
	Name    string `toml:"name,omitempty"`
	Enabled *bool  `toml:"enabled,omitempty"`
	Hosts   Hosts  `toml:"hosts,omitempty"`
}

type Proxies []Proxy

func (proxies Proxies) Len() int           { return len(proxies) }
func (proxies Proxies) Less(i, j int) bool { return proxies[i].Name < proxies[j].Name }
func (proxies Proxies) Swap(i, j int)      { proxies[i], proxies[j] = proxies[j], proxies[i] }

func (proxies *Proxies) Merge(src Proxies) {
	if proxies == nil || len(src) == 0 {
		return
	}

	copied := *proxies
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
	*proxies = copied[:shift+1]
}
