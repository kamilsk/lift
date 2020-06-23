package cnf_test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/kamilsk/lift/internal/cnf"
)

func TestGoModule(t *testing.T) {
	fs := afero.NewMemMapFs()

	t.Run("normal case", func(t *testing.T) {
		require.NoError(t, fs.Mkdir("testdata", 0755))
		require.NoError(t, afero.WriteFile(fs, "testdata/go.mod", []byte("\tmodule \tgo/module/test"), 0644))

		mod, is := GoModule(WorkDir{FS: fs, Path: "testdata"})
		assert.True(t, is)
		assert.Equal(t, "go/module/test", mod)
	})

	t.Run("no exists", func(t *testing.T) {
		mod, is := GoModule(WorkDir{FS: fs, Path: "no-exists"})
		assert.False(t, is)
		assert.Empty(t, mod)
	})

	t.Run("empty file", func(t *testing.T) {
		require.NoError(t, fs.Mkdir("testdata/empty", 0755))
		require.NoError(t, afero.WriteFile(fs, "testdata/empty/go.mod", []byte{}, 0644))

		mod, is := GoModule(WorkDir{FS: fs, Path: "testdata/empty"})
		assert.False(t, is)
		assert.Empty(t, mod)
	})

	t.Run("issue 21", func(t *testing.T) {
		require.NoError(t, fs.Mkdir("testdata/issue-21", 0755))
		require.NoError(t, afero.WriteFile(fs,
			"testdata/issue-21/go.mod",
			[]byte("module go.avito.ru/swat/service-swaha"), 0644))

		mod, is := GoModule(WorkDir{FS: fs, Path: "testdata/issue-21"})
		assert.True(t, is)
		assert.Equal(t, "go.avito.ru/swat/service-swaha", mod)
	})
}

func TestGoPackage(t *testing.T) {
	fs := afero.NewMemMapFs()

	t.Run("normal case", func(t *testing.T) {
		require.NoError(t, fs.MkdirAll("testdata/src/go/package/test", 0755))
		require.NoError(t, afero.WriteFile(fs, "testdata/src/go/package/test/Gopkg.toml", []byte{}, 0644))

		pkg, is := GoPackage(WorkDir{FS: fs, Path: "testdata/src/go/package/test"})
		assert.True(t, is)
		assert.Equal(t, "go/package/test", pkg)
	})

	t.Run("no exists", func(t *testing.T) {
		pkg, is := GoPackage(WorkDir{FS: fs, Path: "no-exists"})
		assert.False(t, is)
		assert.Empty(t, pkg)
	})
}
