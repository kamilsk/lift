package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestSFTP_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *SFTP
		assert.NotPanics(t, func() { dst.Merge(&SFTP{Enabled: pointer.ToBool(true)}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(SFTP)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := SFTP{
			Enabled: pointer.ToBool(false),
			Size:    "small",
		}
		src := SFTP{
			Enabled: pointer.ToBool(false),
			Size:    "medium",
		}

		dst.Merge(&src)
		assert.Equal(t, SFTP{
			Enabled: pointer.ToBool(false),
			Size:    "medium",
		}, dst)
	})
}
