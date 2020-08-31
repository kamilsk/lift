package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestProxies_Merge(t *testing.T) {
	t.Run("nil proxies", func(t *testing.T) {
		var proxies *Proxies
		assert.NotPanics(t, func() { proxies.Merge(Proxies{{Name: "test"}}) })
		assert.Nil(t, proxies)
	})
}

func TestProxies_Sort(t *testing.T) {
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
}
