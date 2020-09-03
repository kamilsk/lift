package model

// A EnvironmentVariables contains key-value pairs of environment variables.
type EnvironmentVariables map[string]string

// Merge combines two set of environment variables.
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
