package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestQueues_Merge(t *testing.T) {
	t.Run("nil queues", func(t *testing.T) {
		var queues *Queues
		assert.NotPanics(t, func() { queues.Merge(Queues{{Name: "test"}}) })
		assert.Nil(t, queues)
	})
}
