package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestElasticSearch_Merge(t *testing.T) {
	t.Run("nil elasticsearch", func(t *testing.T) {
		var elastic *ElasticSearch
		assert.NotPanics(t, func() { elastic.Merge(&ElasticSearch{Version: "6.3.0"}) })
		assert.Nil(t, elastic)
	})
}
