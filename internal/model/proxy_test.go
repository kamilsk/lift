package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestProxies_Merge(t *testing.T) {
	t.Run("nil proxies", func(t *testing.T) {
		var proxies *Proxies
		assert.NotPanics(t, func() { proxies.Merge(Proxies{{Name: "test"}}) })
		assert.Nil(t, proxies)
	})
}
