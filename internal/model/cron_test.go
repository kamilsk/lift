package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestCrons_Merge(t *testing.T) {
	t.Run("nil crons", func(t *testing.T) {
		var crons *Crons
		assert.NotPanics(t, func() { crons.Merge(Crons{{Name: "test"}}) })
		assert.Nil(t, crons)
	})
}

func TestCrons_Sort(t *testing.T) {
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
}
