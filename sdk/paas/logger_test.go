package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestLogger_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Logger
		assert.NotPanics(t, func() { dst.Merge(&Logger{Level: "error"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Logger)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Logger{
			Level: "info",
		}
		src := Logger{
			Level: "error",
		}

		dst.Merge(&src)
		assert.Equal(t, Logger{
			Level: "error",
		}, dst)
	})
}
