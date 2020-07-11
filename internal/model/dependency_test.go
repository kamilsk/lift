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

func TestDependencies_Serialization(t *testing.T) {
	var specification struct {
		Deps Dependencies `toml:"dependencies,omitempty"`
		Envs map[string]struct {
			Deps Dependencies `toml:"dependencies,omitempty"`
		} `toml:"envs,omitempty"`
	}

	tree, err := toml.LoadFile("./testdata/dependencies.toml")
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

func TestDependencies_Sorting(t *testing.T) {
	tests := map[string]struct {
		input    Dependencies
		expected Dependencies
	}{
		"sorted": {
			input: Dependencies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "d"},
			},
			expected: Dependencies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "d"},
			},
		},
		"unsorted": {
			input: Dependencies{
				{Name: "d"},
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Dependencies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
				{Name: "d"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sort.Sort(test.input)
			assert.Equal(t, test.expected, test.input)
		})
	}
}
