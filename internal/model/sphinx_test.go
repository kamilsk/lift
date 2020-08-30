package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSphinxes_Merge(t *testing.T) {
	t.Run("nil sphinxes", func(t *testing.T) {
		var sphinxes *Sphinxes
		assert.NotPanics(t, func() { sphinxes.Merge(Sphinxes{{Name: "test"}}) })
		assert.Nil(t, sphinxes)
	})
}
