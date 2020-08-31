package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestWorkers_Merge(t *testing.T) {
	t.Run("nil workers", func(t *testing.T) {
		var workers *Workers
		assert.NotPanics(t, func() { workers.Merge(Workers{{Name: "test"}}) })
		assert.Nil(t, workers)
	})
}

func TestWorkers_Sort(t *testing.T) {
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
}
