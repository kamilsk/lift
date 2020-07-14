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

func TestMerge(t *testing.T) {
	var spec1, spec2 Application
	{
		tree, err := toml.LoadFile("./testdata/env_vars.toml")
		require.NoError(t, err)

		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &spec1,
			TagName: "toml",
		})
		require.NoError(t, err)
		require.NoError(t, decoder.Decode(tree.ToMap()))
	}
	{
		tree, err := toml.LoadFile("./testdata/dependencies.toml")
		require.NoError(t, err)

		decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			Result:  &spec2,
			TagName: "toml",
		})
		require.NoError(t, err)
		require.NoError(t, decoder.Decode(tree.ToMap()))
	}
	spec1.Merge(spec2)

	tree, err := toml.LoadFile("./testdata/app.toml")
	require.NoError(t, err)

	buf := bytes.NewBuffer(make([]byte, 0, 1024))
	require.NoError(t, toml.NewEncoder(buf).Encode(spec1))
	assert.Equal(t, tree.String(), buf.String())
}
