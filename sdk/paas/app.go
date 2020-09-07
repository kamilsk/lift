package paas

import "sort"

// An Application contains configuration for a service.
type Application struct {
	Specification `toml:",omitempty,squash"`
	Envs          map[string]*Specification `toml:"envs,omitempty"`
}

// Merge combines two service configurations.
func (app *Application) Merge(apps ...Application) {
	if app == nil || len(apps) == 0 {
		return
	}

	for _, src := range apps {
		app.Specification.Merge(&(src.Specification))
		sort.Sort(app.Specification.Dependencies)

		if app.Envs == nil && len(src.Envs) > 0 {
			app.Envs = make(map[string]*Specification)
		}
		for env, spec := range src.Envs {
			if app.Envs[env] == nil {
				app.Envs[env] = new(Specification)
			}
			app.Envs[env].Merge(spec)
			sort.Sort(app.Envs[env].Dependencies)
		}
	}
}
