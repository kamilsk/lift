package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestDependencies_Merge(t *testing.T) {
	t.Run("nil dependencies", func(t *testing.T) {
		var deps *Dependencies
		assert.NotPanics(t, func() { deps.Merge(Dependencies{{Name: "test"}}) })
		assert.Nil(t, deps)
	})
}

func TestDependencies_Sort(t *testing.T) {
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
}
