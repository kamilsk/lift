package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestEnvironmentVariables_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *EnvironmentVariables
		assert.NotPanics(t, func() { dst.Merge(EnvironmentVariables{"KEY": "value"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(EnvironmentVariables)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := EnvironmentVariables{
			"KEY_A": "value",
		}
		src := EnvironmentVariables{
			"KEY_B": "value",
		}

		dst.Merge(src)
		assert.Equal(t, EnvironmentVariables{
			"KEY_A": "value",
			"KEY_B": "value",
		}, dst)
	})

	t.Run("nil map", func(t *testing.T) {
		var dst EnvironmentVariables
		src := EnvironmentVariables{"KEY": "value"}

		assert.NotPanics(t, func() { dst.Merge(src) })
		assert.Equal(t, EnvironmentVariables{"KEY": "value"}, dst)
	})
}
