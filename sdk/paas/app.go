package paas

import "sort"

// An Application contains configuration for a service.
type Application struct {
	Specification `toml:",omitempty,squash"`
	Envs          map[string]*Specification `toml:"envs,omitempty"`
}

// Merge combines two service configurations.
func (dst *Application) Merge(sources ...Application) {
	if dst == nil || len(sources) == 0 {
		return
	}

	for _, src := range sources {
		dst.Specification.Merge(&(src.Specification))

		if dst.Envs == nil && len(src.Envs) > 0 {
			dst.Envs = make(map[string]*Specification)
		}
		for env, spec := range src.Envs {
			if dst.Envs[env] == nil {
				dst.Envs[env] = new(Specification)
			}
			dst.Envs[env].Merge(spec)
		}
	}

	if len(dst.Specification.Dependencies) > 0 && !sort.IsSorted(dst.Specification.Dependencies) {
		sort.Sort(dst.Specification.Dependencies)
	}
	for _, spec := range dst.Envs {
		if len(spec.Dependencies) > 0 && !sort.IsSorted(spec.Dependencies) {
			sort.Sort(spec.Dependencies)
		}
	}
}
