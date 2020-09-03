package model

// A Resource contains configuration for a unit of resource.
type Resource struct {
	CPU    uint `toml:"cpu,omitempty"`
	Memory uint `toml:"memory,omitempty"`
}

// Merge combines two unit of resource configurations.
func (dst *Resource) Merge(src *Resource) {
	if dst == nil || src == nil {
		return
	}

	if src.CPU != 0 {
		dst.CPU = src.CPU
	}
	if src.Memory != 0 {
		dst.Memory = src.Memory
	}
}

// A Resources contains configuration for a Kubernetes resource.
type Resources struct {
	Requests *Resource `toml:"requests,omitempty"`
	Limits   *Resource `toml:"limits,omitempty"`
}

// Merge combines two Kubernetes resource configurations.
func (dst *Resources) Merge(src *Resources) {
	if dst == nil || src == nil {
		return
	}

	if src.Requests != nil && dst.Requests == nil {
		dst.Requests = new(Resource)
	}
	dst.Requests.Merge(src.Requests)

	if src.Limits != nil && dst.Limits == nil {
		dst.Limits = new(Resource)
	}
	dst.Limits.Merge(src.Limits)
}
