package model

type Resource struct {
	CPU    uint `toml:"cpu,omitempty"`
	Memory uint `toml:"memory,omitempty"`
}

func (resource *Resource) Merge(src *Resource) {
	if resource == nil || src == nil {
		return
	}

	if src.CPU > 0 {
		resource.CPU = src.CPU
	}
	if src.Memory > 0 {
		resource.Memory = src.Memory
	}
}

type Resources struct {
	Requests *Resource `toml:"requests,omitempty"`
	Limits   *Resource `toml:"limits,omitempty"`
}

func (resources *Resources) Merge(src *Resources) {
	if resources == nil || src == nil {
		return
	}

	if src.Requests != nil && resources.Requests == nil {
		resources.Requests = new(Resource)
	}
	resources.Requests.Merge(src.Requests)

	if src.Limits != nil && resources.Limits == nil {
		resources.Limits = new(Resource)
	}
	resources.Limits.Merge(src.Limits)
}
