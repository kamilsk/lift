package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestShards_Merge(t *testing.T) {
	t.Run("nil shards", func(t *testing.T) {
		var shards *Shards
		assert.NotPanics(t, func() { shards.Merge(Shards{{Master: "test"}}) })
		assert.Nil(t, shards)
	})
}

func TestShards_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Shards
		expected Shards
	}{
		"sorted": {
			input: Shards{
				{Master: "a"},
				{Master: "b"},
				{Master: "c"},
			},
			expected: Shards{
				{Master: "a"},
				{Master: "b"},
				{Master: "c"},
			},
		},
		"unsorted": {
			input: Shards{
				{Master: "b"},
				{Master: "c"},
				{Master: "a"},
			},
			expected: Shards{
				{Master: "a"},
				{Master: "b"},
				{Master: "c"},
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
