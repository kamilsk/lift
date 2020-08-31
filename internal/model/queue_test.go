package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestQueues_Merge(t *testing.T) {
	t.Run("nil queues", func(t *testing.T) {
		var queues *Queues
		assert.NotPanics(t, func() { queues.Merge(Queues{{Name: "test"}}) })
		assert.Nil(t, queues)
	})
}

func TestQueues_Sort(t *testing.T) {
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
}
