package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestDependencies_Merge(t *testing.T) {
	t.Run("nil dependencies", func(t *testing.T) {
		var deps *Dependencies
		assert.NotPanics(t, func() { deps.Merge(Dependencies{{Name: "test"}}) })
		assert.Nil(t, deps)
	})
}
