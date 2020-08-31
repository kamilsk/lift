package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestMongoDB_Merge(t *testing.T) {
	t.Run("nil mongodb", func(t *testing.T) {
		var mongodb *MongoDB
		assert.NotPanics(t, func() { mongodb.Merge(&MongoDB{Version: "4.0"}) })
		assert.Nil(t, mongodb)
	})
}
