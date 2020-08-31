package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestExecutable_Merge(t *testing.T) {
	t.Run("nil executable", func(t *testing.T) {
		var exec *Executable
		assert.NotPanics(t, func() { exec.Merge(Executable{{Name: "test"}}) })
		assert.Nil(t, exec)
	})
}

func TestExecutable_Sort(t *testing.T) {
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
}
