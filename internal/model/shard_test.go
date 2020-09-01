package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestShards_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Shards
		assert.NotPanics(t, func() { dst.Merge(Shards{{Master: "test"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Shards)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicates", func(t *testing.T) {
		dst := Shards{
			{
				Master: "master-a",
				Slaves: []string{"slave-a"},
			},
			{
				Master: "master-c",
				Slaves: []string{"slave-c", "slave-a"},
			},
			{
				Master: "master-b",
				Slaves: []string{"slave-d"},
			},
		}
		src := Shards{
			{
				Master: "master-d",
				Slaves: []string{"slave-a"},
			},
			{
				Master: "master-c",
				Slaves: []string{"slave-b"},
			},
			{
				Master: "master-a",
				Slaves: []string{"slave-d", "slave-a"},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Shards{
			{
				Master: "master-a",
				Slaves: []string{"slave-a", "slave-d"},
			},
			{
				Master: "master-c",
				Slaves: []string{"slave-c", "slave-a", "slave-b"},
			},
			{
				Master: "master-b",
				Slaves: []string{"slave-d"},
			},
			{
				Master: "master-d",
				Slaves: []string{"slave-a"},
			},
		}, dst)
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
