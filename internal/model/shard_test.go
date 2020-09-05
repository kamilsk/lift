package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestShard_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Shard
		assert.NotPanics(t, func() { dst.Merge(Shard{Primary: "shard"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Shard{Primary: "shard-a"}
		assert.NotPanics(t, func() { dst.Merge(Shard{Primary: "shard-b", Reserve: []string{"shard-c"}}) })
		assert.Nil(t, dst.Reserve)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Shard{
			Primary: "shard",
			Reserve: []string{"shard-c", "shard-a"},
		}
		src := Shard{
			Primary: "shard",
			Reserve: []string{"shard-b", "shard-a"},
		}

		dst.Merge(src)
		assert.Equal(t, Shard{
			Primary: "shard",
			Reserve: []string{"shard-c", "shard-a", "shard-b"},
		}, dst)
	})
}

func TestShards_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Shards
		assert.NotPanics(t, func() { dst.Merge(Shards{{Primary: "shard"}}) })
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
				Primary: "shard-a",
				Reserve: []string{"shard-b"},
			},
			{
				Primary: "shard-c",
				Reserve: []string{"shard-b", "shard-a"},
			},
			{
				Primary: "shard-b",
				Reserve: []string{"shard-d"},
			},
		}
		src := Shards{
			{
				Primary: "shard-d",
				Reserve: []string{"shard-a"},
			},
			{
				Primary: "shard-c",
				Reserve: []string{"shard-b"},
			},
			{
				Primary: "shard-a",
				Reserve: []string{"shard-d", "shard-b"},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Shards{
			{
				Primary: "shard-a",
				Reserve: []string{"shard-b", "shard-d"},
			},
			{
				Primary: "shard-c",
				Reserve: []string{"shard-b", "shard-a"},
			},
			{
				Primary: "shard-b",
				Reserve: []string{"shard-d"},
			},
			{
				Primary: "shard-d",
				Reserve: []string{"shard-a"},
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
				{Primary: "a"},
				{Primary: "b"},
				{Primary: "c"},
			},
			expected: Shards{
				{Primary: "a"},
				{Primary: "b"},
				{Primary: "c"},
			},
		},
		"unsorted": {
			input: Shards{
				{Primary: "b"},
				{Primary: "c"},
				{Primary: "a"},
			},
			expected: Shards{
				{Primary: "a"},
				{Primary: "b"},
				{Primary: "c"},
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
