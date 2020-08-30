package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestEngine_Merge(t *testing.T) {
	t.Run("nil engine", func(t *testing.T) {
		var engine *Engine
		assert.NotPanics(t, func() { engine.Merge(&Engine{Name: "test"}) })
		assert.Nil(t, engine)
	})
}
