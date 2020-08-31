package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestHosts_Merge(t *testing.T) {
	t.Run("nil hosts", func(t *testing.T) {
		var hosts *Hosts
		assert.NotPanics(t, func() { hosts.Merge(Hosts{{Name: "test"}}) })
		assert.Nil(t, hosts)
	})
}

func TestHosts_Sort(t *testing.T) {
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
}
