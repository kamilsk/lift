package model_test

import (
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/require"

	. "github.com/kamilsk/lift/internal/model"
)

func TestEnvironmentVariables_Serialization(t *testing.T) {
	type document struct {
		Base EnvironmentVariables `toml:"env_vars,omitempty"`
		Perf EnvironmentVariables `toml:"envs.perf.env_vars,omitempty"`
		Prod EnvironmentVariables `toml:"envs.prod.env_vars,omitempty"`
	}
	var specification document

	config := new(mapstructure.DecoderConfig)
	config.DecodeHook = Environment()
	config.Result = &specification
	config.TagName = "toml"

	tree, err := toml.LoadFile("./testdata/env_vars.toml")
	require.NoError(t, err)

	decoder, err := mapstructure.NewDecoder(config)
	require.NoError(t, err)
	require.NoError(t, decoder.Decode(tree.ToMap()))
}
