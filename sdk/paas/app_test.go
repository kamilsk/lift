package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestApplication_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Application
		assert.NotPanics(t, func() { dst.Merge(Application{Specification: Specification{Name: "service"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Application)
		assert.NotPanics(t, func() { dst.Merge() })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Application{
			Specification: Specification{
				Dependencies: Dependencies{
					{Name: "service-b"},
					{Name: "service-d"},
				},
			},
		}
		src := Application{
			Specification: Specification{
				Name: "service",
				Dependencies: Dependencies{
					{Name: "service-c"},
					{Name: "service-a"},
				},
			},
			Envs: map[string]*Specification{
				"local": {
					Name: "service",
					Dependencies: Dependencies{
						{Name: "service-d"},
						{Name: "service-a"},
						{Name: "service-c"},
						{Name: "service-b"},
					},
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Application{
			Specification: Specification{
				Name: "service",
				Dependencies: Dependencies{
					{Name: "service-a"},
					{Name: "service-b"},
					{Name: "service-c"},
					{Name: "service-d"},
				},
			},
			Envs: map[string]*Specification{
				"local": {
					Name: "service",
					Dependencies: Dependencies{
						{Name: "service-a"},
						{Name: "service-b"},
						{Name: "service-c"},
						{Name: "service-d"},
					},
				},
			},
		}, dst)
	})
}
