package model

func (app *Application) Merge(apps ...Application) {
	for _, src := range apps {
		app.Specification.Merge(&src.Specification)
		for env, spec := range src.Envs {
			app.Envs[env].Merge(spec)
		}
	}
}

func (spec *Specification) Merge(src *Specification) {
	if spec == nil || src == nil {
		return
	}
	if src.Name != "" {
		spec.Name = src.Name
	}
	if src.Description != "" {
		spec.Description = src.Description
	}
	if src.Kind != "" {
		spec.Kind = src.Kind
	}
	if src.Host != "" {
		spec.Kind = src.Kind
	}
	if src.Replicas > 0 {
		spec.Replicas = src.Replicas
	}
	spec.Engine.Merge(src.Engine)
	spec.Logger.Merge(src.Logger)
	spec.Balancing.Merge(src.Balancing)
	spec.SFTP.Merge(src.SFTP)
}

func (engine *Engine) Merge(src *Engine) {
	if engine == nil || src == nil {
		return
	}
	if src.Name != "" {
		engine.Name = src.Name
	}
	if src.Version != "" {
		engine.Version = src.Version
	}
	if src.Size != "" {
		engine.Size = src.Size
	}
	engine.Resources.Merge(src.Resources)
}

func (logger *Logger) Merge(src *Logger) {
	if logger == nil || src == nil {
		return
	}
	if src.Level != "" {
		logger.Level = src.Level
	}
}

func (balancing *Balancing) Merge(src *Balancing) {
	if balancing == nil || src == nil {
		return
	}
	if src.CookieAffinity != "" {
		balancing.CookieAffinity = src.CookieAffinity
	}
}

func (sftp *SFTP) Merge(src *SFTP) {
	if sftp == nil || src == nil {
		return
	}
	if src.Size != "" {
		sftp.Size = src.Size
	}
	if src.Enabled != nil {
		sftp.Enabled = src.Enabled
	}
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

func (resources *Resources) Merge(src *Resources) {
	if resources == nil || src == nil {
		return
	}
	resources.Requests.Merge(src.Requests)
	resources.Limits.Merge(src.Limits)
}

func (crons Crons) Merge(src Crons) {
	if len(src) == 0 {
		return
	}
}
