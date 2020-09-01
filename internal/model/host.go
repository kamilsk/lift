package model

import "sort"

type Host struct {
	Name        string `toml:"host"`
	AgentPort   uint   `toml:"agent_port"`
	Connections uint   `toml:"connections,omitempty"`
	MaxConns    uint   `toml:"maxconn,omitempty"`
	Weight      uint   `toml:"weight,omitempty"`
	Backup      *bool  `toml:"backup,omitempty"`
}

type Hosts []Host

// Len, Less, Swap implements the sort.Interface.
func (dst Hosts) Len() int           { return len(dst) }
func (dst Hosts) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Hosts) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

func (dst *Hosts) Merge(src Hosts) {
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
