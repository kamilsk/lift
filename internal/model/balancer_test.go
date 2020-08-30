package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestBalancer_Merge(t *testing.T) {
	t.Run("nil balancer", func(t *testing.T) {
		var balancer *Balancer
		assert.NotPanics(t, func() { balancer.Merge(&Balancer{CookieAffinity: "u"}) })
		assert.Nil(t, balancer)
	})
}
