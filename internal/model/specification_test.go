package model_test

import (
	"bytes"
	"sort"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	. "github.com/kamilsk/lift/internal/model"
)

func TestDependencies(t *testing.T) {
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

func TestSorting(t *testing.T) {
	t.Run("Crons", func(t *testing.T) {
		tests := map[string]struct {
			input    Crons
			expected Crons
		}{
			"sorted": {
				input: Crons{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Crons{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Crons{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Crons{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Dependencies", func(t *testing.T) {
		tests := map[string]struct {
			input    Dependencies
			expected Dependencies
		}{
			"sorted": {
				input: Dependencies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Dependencies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Dependencies{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Dependencies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Executable", func(t *testing.T) {
		tests := map[string]struct {
			input    Executable
			expected Executable
		}{
			"sorted": {
				input: Executable{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Executable{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Executable{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Executable{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Hosts", func(t *testing.T) {
		tests := map[string]struct {
			input    Hosts
			expected Hosts
		}{
			"sorted": {
				input: Hosts{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Hosts{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Hosts{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Hosts{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Proxies", func(t *testing.T) {
		tests := map[string]struct {
			input    Proxies
			expected Proxies
		}{
			"sorted": {
				input: Proxies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Proxies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Proxies{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Proxies{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Queues", func(t *testing.T) {
		tests := map[string]struct {
			input    Queues
			expected Queues
		}{
			"sorted": {
				input: Queues{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Queues{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Queues{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Queues{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Sphinxes", func(t *testing.T) {
		tests := map[string]struct {
			input    Sphinxes
			expected Sphinxes
		}{
			"sorted": {
				input: Sphinxes{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Sphinxes{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Sphinxes{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Sphinxes{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
	t.Run("Workers", func(t *testing.T) {
		tests := map[string]struct {
			input    Workers
			expected Workers
		}{
			"sorted": {
				input: Workers{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
				expected: Workers{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
			"unsorted": {
				input: Workers{
					{Name: "b"},
					{Name: "c"},
					{Name: "a"},
				},
				expected: Workers{
					{Name: "a"},
					{Name: "b"},
					{Name: "c"},
				},
			},
		}

		for name, test := range tests {
			t.Run(name, func(t *testing.T) {
				sort.Sort(test.input)
				assert.Equal(t, test.expected, test.input)
			})
		}
	})
}
