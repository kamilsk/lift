package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestMongoDB_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *MongoDB
		assert.NotPanics(t, func() { dst.Merge(&MongoDB{Version: "4.0"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(MongoDB)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := MongoDB{
			Enabled: pointer.ToBool(false),
			Version: "3.6",
			Size:    "small",
		}
		src := MongoDB{
			Enabled: pointer.ToBool(true),
			Version: "4.0",
			Size:    "medium",
		}

		dst.Merge(&src)
		assert.Equal(t, MongoDB{
			Enabled: pointer.ToBool(true),
			Version: "4.0",
			Size:    "medium",
		}, dst)
	})
}
