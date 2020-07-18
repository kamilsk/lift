package model

import "sort"

func (app *Application) Merge(apps ...Application) {
	if app == nil {
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

func (balancing *Balancing) Merge(src *Balancing) {
	if balancing == nil || src == nil {
		return
	}

	if src.CookieAffinity != "" {
		balancing.CookieAffinity = src.CookieAffinity
	}
}

func (crons *Crons) Merge(src Crons) {
	if crons == nil || len(src) == 0 {
		return
	}

	copied := *crons
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*crons = copied[:j+1]
}

func (deps *Dependencies) Merge(src Dependencies) {
	if deps == nil || len(src) == 0 {
		return
	}

	copied := *deps
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*deps = copied[:j+1]
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

	if src.Resources != nil && engine.Resources == nil {
		engine.Resources = new(Resources)
	}
	engine.Resources.Merge(src.Resources)
}

func (vars *EnvironmentVariables) Merge(src EnvironmentVariables) {
	if vars == nil || len(src) == 0 {
		return
	}
	if *vars == nil {
		*vars = make(EnvironmentVariables)
	}
	dst := *vars
	for env, val := range src {
		dst[env] = val
	}
}

func (exec *Executable) Merge(src Executable) {
	if exec == nil || len(src) == 0 {
		return
	}

	copied := *exec
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*exec = copied[:j+1]
}

func (logger *Logger) Merge(src *Logger) {
	if logger == nil || src == nil {
		return
	}

	if src.Level != "" {
		logger.Level = src.Level
	}
}

func (proxies *Proxies) Merge(src Proxies) {
	if proxies == nil || len(src) == 0 {
		return
	}

	copied := *proxies
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*proxies = copied[:j+1]
}

func (queues *Queues) Merge(src Queues) {
	if queues == nil || len(src) == 0 {
		return
	}

	copied := *queues
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*queues = copied[:j+1]
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

	if src.Requests != nil && resources.Requests == nil {
		resources.Requests = new(Resource)
	}
	resources.Requests.Merge(src.Requests)

	if src.Limits != nil && resources.Limits == nil {
		resources.Limits = new(Resource)
	}
	resources.Limits.Merge(src.Limits)
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
		spec.Host = src.Host
	}
	if src.Replicas > 0 {
		spec.Replicas = src.Replicas
	}

	if src.Engine != nil && spec.Engine == nil {
		spec.Engine = new(Engine)
	}
	spec.Engine.Merge(src.Engine)

	if src.Logger != nil && spec.Logger == nil {
		spec.Logger = new(Logger)
	}
	spec.Logger.Merge(src.Logger)

	if src.Balancing != nil && spec.Balancing == nil {
		spec.Balancing = new(Balancing)
	}
	spec.Balancing.Merge(src.Balancing)

	if src.SFTP != nil && spec.SFTP == nil {
		spec.SFTP = new(SFTP)
	}
	spec.SFTP.Merge(src.SFTP)

	spec.Crons.Merge(src.Crons)
	spec.Dependencies.Merge(src.Dependencies)
	spec.Executable.Merge(src.Executable)
	spec.Proxies.Merge(src.Proxies)
	spec.Queues.Merge(src.Queues)
	spec.Workers.Merge(src.Workers)
	spec.EnvVars.Merge(src.EnvVars)
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

func (workers *Workers) Merge(src Workers) {
	if workers == nil || len(src) == 0 {
		return
	}

	copied := *workers
	copied = append(copied, src...)
	sort.Sort(copied)
	j := 0
	for i := 1; i < len(copied); i++ {
		if copied[j].Name == copied[i].Name {
			continue
		}
		j++
		copied[j] = copied[i]
	}
	*workers = copied[:j+1]
}
