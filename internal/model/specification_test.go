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

func TestSpecification_Merge(t *testing.T) {
	t.Run("nil specification", func(t *testing.T) {
		var spec *Specification
		assert.NotPanics(t, func() { spec.Merge(&Specification{Name: "test"}) })
		assert.Nil(t, spec)
	})
}

func TestDependencies(t *testing.T) {
	t.SkipNow()

	var app Application

	tree, err := toml.LoadFile("testdata/dependencies.toml")
	require.NoError(t, err)

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &app,
		TagName: "toml",
	})
	require.NoError(t, err)
	require.NoError(t, decoder.Decode(tree.ToMap()))

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(app))
	assert.Equal(t, tree.String(), buf.String())
}

func TestEnvironmentVariables(t *testing.T) {
	t.SkipNow()

	var app Application

	tree, err := toml.LoadFile("testdata/env_vars.toml")
	require.NoError(t, err)

	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:  &app,
		TagName: "toml",
	})
	require.NoError(t, err)
	require.NoError(t, decoder.Decode(tree.ToMap()))

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(app))
	assert.Equal(t, tree.String(), buf.String())
}
