package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestExecutable_Merge(t *testing.T) {
	t.Run("nil executable", func(t *testing.T) {
		var exec *Executable
		assert.NotPanics(t, func() { exec.Merge(Executable{{Name: "test"}}) })
		assert.Nil(t, exec)
	})
}
