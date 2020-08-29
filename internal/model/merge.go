package model

import (
	"sort"
)

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
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*exec = copied[:shift+1]
}

func (proxies *Proxies) Merge(src Proxies) {
	if proxies == nil || len(src) == 0 {
		return
	}

	copied := *proxies
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*proxies = copied[:shift+1]
}

func (queues *Queues) Merge(src Queues) {
	if queues == nil || len(src) == 0 {
		return
	}

	copied := *queues
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*queues = copied[:shift+1]
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

	if src.Balancer != nil && spec.Balancer == nil {
		spec.Balancer = new(Balancer)
	}
	spec.Balancer.Merge(src.Balancer)

	if src.PostgreSQL != nil && spec.PostgreSQL == nil {
		spec.PostgreSQL = new(PostgreSQL)
	}
	spec.PostgreSQL.Merge(src.PostgreSQL)

	if src.Redis != nil && spec.Redis == nil {
		spec.Redis = new(Redis)
	}
	spec.Redis.Merge(src.Redis)

	if src.RedisSharded != nil && spec.RedisSharded == nil {
		spec.RedisSharded = new(ShardedRedis)
	}
	spec.RedisSharded.Merge(src.RedisSharded)

	if src.SFTP != nil && spec.SFTP == nil {
		spec.SFTP = new(SFTP)
	}
	spec.SFTP.Merge(src.SFTP)

	spec.Crons.Merge(src.Crons)
	spec.Dependencies.Merge(src.Dependencies)
	spec.Executable.Merge(src.Executable)
	spec.Proxies.Merge(src.Proxies)
	spec.Queues.Merge(src.Queues)
	spec.Sphinxes.Merge(src.Sphinxes)
	spec.Workers.Merge(src.Workers)
	spec.EnvVars.Merge(src.EnvVars)
}

func (sphinxes *Sphinxes) Merge(src Sphinxes) {
	if sphinxes == nil || len(src) == 0 {
		return
	}

	copied := *sphinxes
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*sphinxes = copied[:shift+1]
}

func (workers *Workers) Merge(src Workers) {
	if workers == nil || len(src) == 0 {
		return
	}

	copied := *workers
	copied = append(copied, src...)
	sort.Sort(copied)
	shift := 0
	for i := 1; i < len(copied); i++ {
		if copied[shift].Name == copied[i].Name {
			continue
		}
		shift++
		copied[shift] = copied[i]
	}
	*workers = copied[:shift+1]
}
