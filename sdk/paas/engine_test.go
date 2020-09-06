package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestEngine_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Engine
		assert.NotPanics(t, func() { dst.Merge(&Engine{Name: "golang"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Engine)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Engine{
			Name:    "golang",
			Version: "1.11",
			Size:    "small",
		}
		src := Engine{
			Name:    "golang",
			Version: "9.6",
			Size:    "medium",
			Resources: &Resources{
				Requests: &Resource{
					CPU:    1,
					Memory: 100,
				},
				Limits: &Resource{
					CPU:    2,
					Memory: 200,
				},
			},
		}

		dst.Merge(&src)
		assert.Equal(t, Engine{
			Name:    "golang",
			Version: "9.6",
			Size:    "medium",
			Resources: &Resources{
				Requests: &Resource{
					CPU:    1,
					Memory: 100,
				},
				Limits: &Resource{
					CPU:    2,
					Memory: 200,
				},
			},
		}, dst)
	})
}
