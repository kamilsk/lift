package model

import "sort"

type Host struct {
	Name        string `toml:"host,omitempty"`
	AgentPort   uint   `toml:"agent_port,omitempty"`
	Connections uint   `toml:"connections,omitempty"`
	MaxConns    uint   `toml:"maxconn,omitempty"`
	Weight      uint   `toml:"weight,omitempty"`
	Backup      bool   `toml:"backup,omitempty"`
}

type Hosts []Host

func (hosts Hosts) Len() int           { return len(hosts) }
func (hosts Hosts) Less(i, j int) bool { return hosts[i].Name < hosts[j].Name }
func (hosts Hosts) Swap(i, j int)      { hosts[i], hosts[j] = hosts[j], hosts[i] }

func (hosts *Hosts) Merge(src Hosts) {
	if hosts == nil || len(src) == 0 {
		return
	}

	copied := *hosts
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
	*hosts = copied[:shift+1]
}
