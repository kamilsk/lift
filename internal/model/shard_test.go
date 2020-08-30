package model_test

import (
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
