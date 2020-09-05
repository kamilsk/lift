package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestRedis_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Redis
		assert.NotPanics(t, func() { dst.Merge(&Redis{Version: "5.0"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Redis)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Redis{
			Enabled:  pointer.ToBool(false),
			Version:  "4.0",
			Size:     "small",
			Type:     "data",
			Replicas: 1,
		}
		src := Redis{
			Enabled:  pointer.ToBool(true),
			Version:  "5.0",
			Size:     "medium",
			Type:     "cache",
			Replicas: 10,
		}

		dst.Merge(&src)
		assert.Equal(t, Redis{
			Enabled:  pointer.ToBool(true),
			Version:  "5.0",
			Size:     "medium",
			Type:     "cache",
			Replicas: 10,
		}, dst)
	})
}

func TestShardedRedis_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *ShardedRedis
		assert.NotPanics(t, func() { dst.Merge(&ShardedRedis{Version: "5.0"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(ShardedRedis)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := ShardedRedis{
			Enabled:     pointer.ToBool(false),
			Version:     "4.0",
			Size:        "small",
			Shards:      nil,
			SelfSharded: pointer.ToBool(true),
		}
		src := ShardedRedis{
			Enabled: pointer.ToBool(true),
			Version: "5.0",
			Size:    "medium",
			Shards: Shards{
				{
					Primary: "master-a",
					Reserve: []string{"slave-a", "slave-b"},
				},
			},
			SelfSharded: pointer.ToBool(false),
		}

		dst.Merge(&src)
		assert.Equal(t, ShardedRedis{
			Enabled: pointer.ToBool(true),
			Version: "5.0",
			Size:    "medium",
			Shards: Shards{
				{
					Primary: "master-a",
					Reserve: []string{"slave-a", "slave-b"},
				},
			},
			SelfSharded: pointer.ToBool(false),
		}, dst)
	})
}
