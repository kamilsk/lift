package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestPostgreSQL_Merge(t *testing.T) {
	t.Run("nil postgresql", func(t *testing.T) {
		var postgresql *PostgreSQL
		assert.NotPanics(t, func() { postgresql.Merge(&PostgreSQL{Version: "9.6"}) })
		assert.Nil(t, postgresql)
	})
}
