package cmd_test

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/kamilsk/lift/sdk/paas"
)

var update = flag.Bool("update", false, "update golden files")

func TestMerge(t *testing.T) {
	t.SkipNow()

	t.Run("workflow", func(t *testing.T) {
		matches, err := filepath.Glob("testdata/components/*/*.toml")
		require.NoError(t, err)

		specs := make([]Application, 0, len(matches))
		for _, path := range matches {
			tree, err := toml.LoadFile(path)
			require.NoError(t, err)

			var spec Application
			decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
				Result:  &spec,
				TagName: "toml",
			})
			require.NoError(t, err)
			require.NoError(t, decoder.Decode(tree.ToMap()))
			specs = append(specs, spec)
		}

		if *update {
			file, err := os.Create("testdata/app.toml")
			require.NoError(t, err)

			app := new(Application)
			app.Merge(specs...)
			require.NoError(t, toml.NewEncoder(file).Encode(app))
			require.NoError(t, file.Close())
		}

		tree, err := toml.LoadFile("testdata/app.toml")
		require.NoError(t, err)

		app := new(Application)
		app.Merge(specs...)
		buf := bytes.NewBuffer(make([]byte, 0, 1024))
		require.NoError(t, toml.NewEncoder(buf).Encode(app))
		assert.Equal(t, tree.String(), buf.String())
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
