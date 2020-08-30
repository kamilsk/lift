package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestEnvironmentVariables_Merge(t *testing.T) {
	t.Run("nil env vars", func(t *testing.T) {
		var vars *EnvironmentVariables
		assert.NotPanics(t, func() { vars.Merge(EnvironmentVariables{"ENV": "test"}) })
		assert.Nil(t, vars)
	})
}
