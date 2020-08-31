package model

type Application struct {
	Specification `toml:",omitempty,squash"`
	Envs          map[string]*Specification `toml:"envs,omitempty"`
}

func (app *Application) Merge(apps ...Application) {
	if app == nil || len(apps) == 0 {
		return
	}

	if app.Envs == nil {
		app.Envs = make(map[string]*Specification)
	}

	for _, src := range apps {
		app.Specification.Merge(&(src.Specification))
		for env, spec := range src.Envs {
			if app.Envs[env] == nil {
				app.Envs[env] = new(Specification)
			}
			app.Envs[env].Merge(spec)
		}
	}
}
