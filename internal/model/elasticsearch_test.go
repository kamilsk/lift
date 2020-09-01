package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestElasticSearch_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *ElasticSearch
		assert.NotPanics(t, func() { dst.Merge(&ElasticSearch{Version: "6.3.0"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(ElasticSearch)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := ElasticSearch{
			Enabled: nil,
			Version: "6.2.0",
			Size:    "small",
		}
		src := ElasticSearch{
			Enabled: pointer.ToBool(true),
			Version: "6.3.0",
			Size:    "medium",
		}

		dst.Merge(&src)
		assert.Equal(t, ElasticSearch{
			Enabled: pointer.ToBool(true),
			Version: "6.3.0",
			Size:    "medium",
		}, dst)
	})
}
