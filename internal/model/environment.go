package model

import "sort"

type EnvironmentVariable struct {
	Name  string
	Value string
}

func (env EnvironmentVariable) ToMap() map[string]interface{} {
	return map[string]interface{}{env.Name: env.Value}
}

type EnvironmentVariables []EnvironmentVariable

func (vars EnvironmentVariables) ToMap() map[string]interface{} {
	if len(vars) == 0 {
		return nil
	}
	sorted := make(EnvironmentVariables, len(vars))
	copy(sorted, vars)
	sort.Sort(sorted)
	out := make(map[string]interface{}, len(sorted))
	for _, variable := range sorted {
		out[variable.Name] = variable.Value
	}
	return out
}

func (vars EnvironmentVariables) Len() int           { return len(vars) }
func (vars EnvironmentVariables) Less(i, j int) bool { return vars[i].Name < vars[j].Name }
func (vars EnvironmentVariables) Swap(i, j int)      { vars[i], vars[j] = vars[j], vars[i] }

type EnvironmentWithVariables map[string]EnvironmentVariables

func (sections EnvironmentWithVariables) ToMap() map[string]interface{} {
	if len(sections) == 0 {
		return nil
	}
	out := make(map[string]interface{}, len(sections))
	for env, vars := range sections {
		out[env] = vars.ToMap()
	}
	return out
}
