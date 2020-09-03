package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestResource_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Resource
		assert.NotPanics(t, func() { dst.Merge(&Resource{CPU: 1}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Resource)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Resource{
			CPU:    1,
			Memory: 10,
		}
		src := Resource{
			CPU:    2,
			Memory: 20,
		}

		dst.Merge(&src)
		assert.Equal(t, Resource{
			CPU:    2,
			Memory: 20,
		}, dst)
	})
}

func TestResources_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Resources
		assert.NotPanics(t, func() { dst.Merge(&Resources{Requests: &Resource{CPU: 1}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Resources)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Resources{
			Requests: &Resource{
				CPU:    1,
				Memory: 10,
			},
		}
		src := Resources{
			Limits: &Resource{
				CPU:    2,
				Memory: 20,
			},
		}

		dst.Merge(&src)
		assert.Equal(t, Resources{
			Requests: &Resource{
				CPU:    1,
				Memory: 10,
			},
			Limits: &Resource{
				CPU:    2,
				Memory: 20,
			},
		}, dst)
	})
}
