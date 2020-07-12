package model_test

import (
	"bytes"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/kamilsk/lift/internal/model"
)

func TestEnvironmentVariables_Serialization(t *testing.T) {
	type environment struct {
		Vars EnvironmentVariables `toml:"env_vars,omitempty"`
	}

	var specification struct {
		Vars EnvironmentVariables   `toml:"env_vars,omitempty"`
		Envs map[string]environment `toml:"envs,omitempty"`
	}

	tree, err := toml.LoadFile("./testdata/env_vars.toml")
	require.NoError(t, err)

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &specification,
		TagName: "toml",
	})
	require.NoError(t, err)
	require.NoError(t, decoder.Decode(tree.ToMap()))

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(specification))
	assert.Equal(t, tree.String(), buf.String())
}

func TestEnvironmentVariables_Sorting(t *testing.T) {
	type environment struct {
		Vars EnvironmentVariables `toml:"env_vars,omitempty"`
	}

	var specification struct {
		Vars EnvironmentVariables   `toml:"env_vars,omitempty"`
		Envs map[string]environment `toml:"envs,omitempty"`
	}
	specification.Vars = EnvironmentVariables{
		"SERVICE_X_TIMEOUT":            "1s",
		"SERVICE_X_MAX_IDLE_CONNS":     "10",
		"SERVICE_X_CONNECTION_TIMEOUT": "1s",
	}
	specification.Envs = map[string]environment{
		"perf": {
			Vars: EnvironmentVariables{
				"SERVICE_X_TIMEOUT":            "300ms",
				"SERVICE_X_MAX_IDLE_CONNS":     "10",
				"SERVICE_X_CONNECTION_TIMEOUT": "100ms",
			},
		},
		"prod": {
			Vars: EnvironmentVariables{
				"SERVICE_X_TIMEOUT":            "300ms",
				"SERVICE_X_MAX_IDLE_CONNS":     "100",
				"SERVICE_X_CONNECTION_TIMEOUT": "100ms",
			},
		},
	}

	tree, err := toml.LoadFile("./testdata/env_vars.toml")
	require.NoError(t, err)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(specification))
	assert.Equal(t, tree.String(), buf.String())
}
