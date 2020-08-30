package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestResource_Merge(t *testing.T) {
	t.Run("nil resource", func(t *testing.T) {
		var resource *Resource
		assert.NotPanics(t, func() { resource.Merge(&Resource{CPU: 1}) })
		assert.Nil(t, resource)
	})
}

func TestResources_Merge(t *testing.T) {
	t.Run("nil resources", func(t *testing.T) {
		var resources *Resources
		assert.NotPanics(t, func() { resources.Merge(&Resources{Requests: &Resource{CPU: 1}}) })
		assert.Nil(t, resources)
	})
}
