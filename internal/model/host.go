package model

// A Host contains configuration for a host.
type Host struct {
	Name        string `toml:"host"`
	AgentPort   uint   `toml:"agent_port,omitempty"`
	Connections uint   `toml:"connections,omitempty"`
	MaxConns    uint   `toml:"maxconn,omitempty"`
	Weight      uint   `toml:"weight,omitempty"`
	Backup      *bool  `toml:"backup,omitempty"`
}

// Merge combines two host configurations.
func (dst *Host) Merge(src Host) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.AgentPort != 0 {
		dst.AgentPort = src.AgentPort
	}
	if src.Connections != 0 {
		dst.Connections = src.Connections
	}
	if src.MaxConns != 0 {
		dst.MaxConns = src.MaxConns
	}
	if src.Weight != 0 {
		dst.Weight = src.Weight
	}
	if src.Backup != nil {
		dst.Backup = src.Backup
	}
}

// Hosts is a list of Host.
type Hosts []Host

// Len, Less, Swap implements the sort.Interface.
func (dst Hosts) Len() int           { return len(dst) }
func (dst Hosts) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Hosts) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of host configurations.
func (dst *Hosts) Merge(src Hosts) {
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
	for i, host := range copied {
		origin := registry[host.Name]
		if i == origin {
			unique = append(unique, host)
			continue
		}
		unique[origin].Merge(host)
	}

	*dst = unique
}
