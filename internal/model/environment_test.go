package model_test

import (
	"bytes"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEnvironments_Serialization(t *testing.T) {
	var specification struct {
		Vars map[string]string `toml:"env_vars,omitempty"`
		Envs map[string]struct {
			Vars map[string]string `toml:"env_vars,omitempty"`
		} `toml:"envs,omitempty"`
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
