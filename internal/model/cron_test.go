package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestCrons_Merge(t *testing.T) {
	t.Run("nil crons", func(t *testing.T) {
		var crons *Crons
		assert.NotPanics(t, func() { crons.Merge(Crons{{Name: "test"}}) })
		assert.Nil(t, crons)
	})
}
