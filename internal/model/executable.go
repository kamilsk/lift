package model

// An Exec contains configuration for unit of work.
type Exec struct {
	Name          string     `toml:"name"`
	Enabled       *bool      `toml:"enabled,omitempty"`
	Command       string     `toml:"command"`
	Image         string     `toml:"image,omitempty"`
	Replicas      uint       `toml:"replicas"`
	Port          uint       `toml:"service-port"`
	RedinessProbe string     `toml:"readiness-probe-command"`
	LivenessProbe string     `toml:"liveness-probe-command"`
	Size          string     `toml:"size"`
	Resources     *Resources `toml:"resources,omitempty"`
}

// Merge combines two unit of work configurations.
func (dst *Exec) Merge(src Exec) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Enabled != nil {
		dst.Enabled = src.Enabled
	}
	if src.Command != "" {
		dst.Command = src.Command
	}
	if src.Image != "" {
		dst.Image = src.Image
	}
	if src.Replicas != 0 {
		dst.Replicas = src.Replicas
	}
	if src.Port != 0 {
		dst.Port = src.Port
	}
	if src.RedinessProbe != "" {
		dst.RedinessProbe = src.RedinessProbe
	}
	if src.LivenessProbe != "" {
		dst.LivenessProbe = src.LivenessProbe
	}
	if src.Size != "" {
		dst.Size = src.Size
	}

	if src.Resources != nil && dst.Resources == nil {
		dst.Resources = new(Resources)
	}
	dst.Resources.Merge(src.Resources)
}

// Executable is a list of Exec.
type Executable []Exec

// Len, Less, Swap implements the sort.Interface.
func (dst Executable) Len() int           { return len(dst) }
func (dst Executable) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Executable) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of unit of work configurations.
func (dst *Executable) Merge(src Executable) {
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
	for i, exec := range copied {
		origin := registry[exec.Name]
		if i == origin {
			unique = append(unique, exec)
			continue
		}
		unique[origin].Merge(exec)
	}

	*dst = unique
}
