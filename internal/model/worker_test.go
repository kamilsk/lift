package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestWorkers_Merge(t *testing.T) {
	t.Run("nil workers", func(t *testing.T) {
		var workers *Workers
		assert.NotPanics(t, func() { workers.Merge(Workers{{Name: "test"}}) })
		assert.Nil(t, workers)
	})
}
