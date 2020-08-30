package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestRedis_Merge(t *testing.T) {
	t.Run("nil redis", func(t *testing.T) {
		var redis *Redis
		assert.NotPanics(t, func() { redis.Merge(&Redis{Version: "5.0"}) })
		assert.Nil(t, redis)
	})
}

func TestShardedRedis_Merge(t *testing.T) {
	t.Run("nil sharded redis", func(t *testing.T) {
		var redis *ShardedRedis
		assert.NotPanics(t, func() { redis.Merge(&ShardedRedis{Version: "5.0"}) })
		assert.Nil(t, redis)
	})
}
