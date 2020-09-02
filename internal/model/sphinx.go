package model

// A Sphinx contains configuration for a database.
type Sphinx struct {
	Name    string `toml:"name,omitepmty"`
	Enabled *bool  `toml:"enabled"`
	Hosts   Hosts  `toml:"hosts"`
	Haproxy string `toml:"haproxy_tag,omitempty"`
}

// Merge combines two database configurations.
func (dst *Sphinx) Merge(src Sphinx) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	dst.Hosts.Merge(src.Hosts)

	if src.Haproxy != "" {
		dst.Haproxy = src.Haproxy
	}
}

// Sphinxes is a list of Sphinx.
type Sphinxes []Sphinx

// Len, Less, Swap implements the sort.Interface.
func (dst Sphinxes) Len() int           { return len(dst) }
func (dst Sphinxes) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Sphinxes) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of database configurations.
func (dst *Sphinxes) Merge(src Sphinxes) {
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
	for i, sphinx := range copied {
		origin := registry[sphinx.Name]
		if i == origin {
			unique = append(unique, sphinx)
			continue
		}
		unique[origin].Merge(sphinx)
	}

	*dst = unique
}
