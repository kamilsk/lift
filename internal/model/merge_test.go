package model_test

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

	. "github.com/kamilsk/lift/internal/model"
)

var update = flag.Bool("update", false, "update golden files")

func TestMerge(t *testing.T) {
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
	t.Run("kernel cases", func(t *testing.T) {
		t.Run("nil application", func(t *testing.T) {
			var app *Application
			assert.NotPanics(t, func() { app.Merge(Application{Specification: Specification{Name: "test"}}) })
			assert.Nil(t, app)
		})
		t.Run("nil balancing", func(t *testing.T) {
			var balancer *Balancing
			assert.NotPanics(t, func() { balancer.Merge(&Balancing{CookieAffinity: "u"}) })
			assert.Nil(t, balancer)
		})
		t.Run("nil crons", func(t *testing.T) {
			var crons *Crons
			assert.NotPanics(t, func() { crons.Merge(Crons{{Name: "test"}}) })
			assert.Nil(t, crons)
		})
		t.Run("nil dependencies", func(t *testing.T) {
			var deps *Dependencies
			assert.NotPanics(t, func() { deps.Merge(Dependencies{{Name: "test"}}) })
			assert.Nil(t, deps)
		})
		t.Run("nil engine", func(t *testing.T) {
			var engine *Engine
			assert.NotPanics(t, func() { engine.Merge(&Engine{Name: "test"}) })
			assert.Nil(t, engine)
		})
		t.Run("nil env vars", func(t *testing.T) {
			var vars *EnvironmentVariables
			assert.NotPanics(t, func() { vars.Merge(EnvironmentVariables{"ENV": "test"}) })
			assert.Nil(t, vars)
		})
		t.Run("nil executable", func(t *testing.T) {
			var exec *Executable
			assert.NotPanics(t, func() { exec.Merge(Executable{{Name: "test"}}) })
			assert.Nil(t, exec)
		})
		t.Run("nil logger", func(t *testing.T) {
			var logger *Logger
			assert.NotPanics(t, func() { logger.Merge(&Logger{Level: "debug"}) })
			assert.Nil(t, logger)
		})
		t.Run("nil proxies", func(t *testing.T) {
			var proxies *Proxies
			assert.NotPanics(t, func() { proxies.Merge(Proxies{{Name: "test"}}) })
			assert.Nil(t, proxies)
		})
		t.Run("nil queues", func(t *testing.T) {
			var queues *Queues
			assert.NotPanics(t, func() { queues.Merge(Queues{{Name: "test"}}) })
			assert.Nil(t, queues)
		})
		t.Run("nil resource", func(t *testing.T) {
			var resource *Resource
			assert.NotPanics(t, func() { resource.Merge(&Resource{CPU: 1}) })
			assert.Nil(t, resource)
		})
		t.Run("nil resources", func(t *testing.T) {
			var resources *Resources
			assert.NotPanics(t, func() { resources.Merge(&Resources{Requests: &Resource{CPU: 1}}) })
			assert.Nil(t, resources)
		})
		t.Run("nil sftp", func(t *testing.T) {
			var sftp *SFTP
			assert.NotPanics(t, func() { sftp.Merge(&SFTP{Size: "small"}) })
			assert.Nil(t, sftp)
		})
		t.Run("nil specification", func(t *testing.T) {
			var spec *Specification
			assert.NotPanics(t, func() { spec.Merge(&Specification{Name: "test"}) })
			assert.Nil(t, spec)
		})
		t.Run("nil sphinxes", func(t *testing.T) {
			var sphinxes *Sphinxes
			assert.NotPanics(t, func() { sphinxes.Merge(Sphinxes{{Name: "test"}}) })
			assert.Nil(t, sphinxes)
		})
		t.Run("nil workers", func(t *testing.T) {
			var workers *Workers
			assert.NotPanics(t, func() { workers.Merge(Workers{{Name: "test"}}) })
			assert.Nil(t, workers)
		})
	})
}
