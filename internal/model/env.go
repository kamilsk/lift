package model

type EnvironmentVariables map[string]string

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
