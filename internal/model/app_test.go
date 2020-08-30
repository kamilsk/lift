package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestApplication_Merge(t *testing.T) {
	t.Run("nil application", func(t *testing.T) {
		var app *Application
		assert.NotPanics(t, func() { app.Merge(Application{Specification: Specification{Name: "test"}}) })
		assert.Nil(t, app)
	})
}
