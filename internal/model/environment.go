package model

type EnvironmentVariable struct {
	Name  string
	Value string
}

func (env EnvironmentVariable) ToMap() map[string]interface{} {
	return map[string]interface{}{env.Name: env.Value}
}

type EnvironmentVariables []EnvironmentVariable

func (env EnvironmentVariables) ToMap() map[string]interface{} {
	out := make(map[string]interface{}, len(env))
	for _, variable := range env {
		out[variable.Name] = variable.Value
	}
	return out
}

func (env EnvironmentVariables) Len() int           { return len(env) }
func (env EnvironmentVariables) Less(i, j int) bool { return env[i].Name < env[j].Name }
func (env EnvironmentVariables) Swap(i, j int)      { env[i], env[j] = env[j], env[i] }
