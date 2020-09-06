package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestBalancer_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Balancer
		assert.NotPanics(t, func() { dst.Merge(&Balancer{CookieAffinity: "u"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Balancer)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Balancer{
			CookieAffinity: "x",
		}
		src := Balancer{
			CookieAffinity: "u",
		}

		dst.Merge(&src)
		assert.Equal(t, Balancer{
			CookieAffinity: "u",
		}, dst)
	})
}
