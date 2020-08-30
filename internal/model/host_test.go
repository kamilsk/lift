package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestHosts_Merge(t *testing.T) {
	t.Run("nil hosts", func(t *testing.T) {
		var hosts *Hosts
		assert.NotPanics(t, func() { hosts.Merge(Hosts{{Name: "test"}}) })
		assert.Nil(t, hosts)
	})
}
