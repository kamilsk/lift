package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSphinxes_Merge(t *testing.T) {
	t.Run("nil sphinxes", func(t *testing.T) {
		var sphinxes *Sphinxes
		assert.NotPanics(t, func() { sphinxes.Merge(Sphinxes{{Name: "test"}}) })
		assert.Nil(t, sphinxes)
	})
}

func TestSphinxes_Sort(t *testing.T) {
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
}
