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

func TestEnvironmentVariable_ToMap(t *testing.T) {
	tests := map[string]struct {
		env      EnvironmentVariable
		expected map[string]interface{}
	}{
		"empty": {expected: map[string]interface{}{"": ""}},
		"filled": {
			env:      EnvironmentVariable{Name: "name", Value: "value"},
			expected: map[string]interface{}{"name": "value"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.env.ToMap())
		})
	}
}

func TestEnvironmentVariables_ToMap(t *testing.T) {
	tests := map[string]struct {
		vars     EnvironmentVariables
		expected map[string]interface{}
	}{
		"empty": {},
		"filled": {
			vars:     EnvironmentVariables{{Name: "name", Value: "value"}},
			expected: map[string]interface{}{"name": "value"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.vars.ToMap())
		})
	}
}

func TestEnvironmentWithVariables_Serialization(t *testing.T) {
	var specification EnvironmentWithVariables

	tree, err := toml.LoadFile("./testdata/env_vars.toml")
	require.NoError(t, err)

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		DecodeHook: Environment(),
		Result:     &specification,
	})
	require.NoError(t, err)
	require.NoError(t, decoder.Decode(tree.ToMap()))

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(specification.ToMap()))
	assert.Equal(t, tree.String(), buf.String())
}

func TestEnvironmentWithVariables_ToMap(t *testing.T) {
	tests := map[string]struct {
		sections EnvironmentWithVariables
		expected map[string]interface{}
	}{
		"empty": {},
		"filled": {
			sections: EnvironmentWithVariables{
				"perf": EnvironmentVariables{{Name: "name", Value: "value"}},
				"prod": EnvironmentVariables{{Name: "name", Value: "value"}},
			},
			expected: map[string]interface{}{
				"perf": map[string]interface{}{"name": "value"},
				"prod": map[string]interface{}{"name": "value"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, test.expected, test.sections.ToMap())
		})
	}
}
