package model

// A Proxy contains configuration for an abstract proxy.
type Proxy struct {
	Name    string `toml:"name,omitempty"`
	Enabled *bool  `toml:"enabled"`
	Hosts   Hosts  `toml:"hosts"`
}

// Merge combines two proxy configurations.
func (dst *Proxy) Merge(src Proxy) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	dst.Hosts.Merge(src.Hosts)
}

// Proxies is a list of Proxy.
type Proxies []Proxy

// Len, Less, Swap implements the sort.Interface.
func (dst Proxies) Len() int           { return len(dst) }
func (dst Proxies) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Proxies) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of proxy configurations.
func (dst *Proxies) Merge(src Proxies) {
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
	for i, proxy := range copied {
		origin := registry[proxy.Name]
		if i == origin {
			unique = append(unique, proxy)
			continue
		}
		unique[origin].Merge(proxy)
	}

	*dst = unique
}
