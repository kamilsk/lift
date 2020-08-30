package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestLogger_Merge(t *testing.T) {
	t.Run("nil logger", func(t *testing.T) {
		var logger *Logger
		assert.NotPanics(t, func() { logger.Merge(&Logger{Level: "debug"}) })
		assert.Nil(t, logger)
	})
}
