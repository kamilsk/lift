package model

// A Dependency describes a service dependency.
type Dependency struct {
	Name         string `toml:"name"`
	Mock         *bool  `toml:"mock,omitempty"`
	MockReplicas uint   `toml:"mock-replicas,omitempty"`
}

// Merge combines two service dependencies.
func (dst *Dependency) Merge(src Dependency) {
	if dst == nil || dst.Name != src.Name {
		return
	}

	if src.Mock != nil {
		dst.Mock = src.Mock
	}
	if src.MockReplicas != 0 {
		dst.MockReplicas = src.MockReplicas
	}
}

// Dependencies is a list of Dependency.
type Dependencies []Dependency

// Len, Less, Swap implements the sort.Interface.
func (dst Dependencies) Len() int           { return len(dst) }
func (dst Dependencies) Less(i, j int) bool { return dst[i].Name < dst[j].Name }
func (dst Dependencies) Swap(i, j int)      { dst[i], dst[j] = dst[j], dst[i] }

// Merge combines two set of service dependencies.
func (dst *Dependencies) Merge(src Dependencies) {
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
	for i, dependency := range copied {
		origin := registry[dependency.Name]
		if i == origin {
			unique = append(unique, dependency)
			continue
		}
		unique[origin].Merge(dependency)
	}

	*dst = unique
}
