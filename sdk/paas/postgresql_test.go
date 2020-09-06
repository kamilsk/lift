package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestPostgreSQL_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *PostgreSQL
		assert.NotPanics(t, func() { dst.Merge(&PostgreSQL{Version: "9.6"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(PostgreSQL)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := PostgreSQL{
			Enabled:  nil,
			Version:  "9.4",
			Size:     "small",
			Fixtures: pointer.ToBool(true),
			OwnName:  pointer.ToBool(true),
			DataBus:  nil,
		}
		src := PostgreSQL{
			Enabled:  pointer.ToBool(true),
			Version:  "9.6",
			Size:     "medium",
			Fixtures: pointer.ToBool(false),
			OwnName:  pointer.ToBool(false),
			DataBus: &DataBus{
				BatchSize: 10,
				Schemas:   []string{"schema-b", "schema-a"},
			},
		}

		dst.Merge(&src)
		assert.Equal(t, PostgreSQL{
			Enabled:  pointer.ToBool(true),
			Version:  "9.6",
			Size:     "medium",
			Fixtures: pointer.ToBool(false),
			OwnName:  pointer.ToBool(false),
			DataBus: &DataBus{
				BatchSize: 10,
				Schemas:   []string{"schema-b", "schema-a"},
			},
		}, dst)
	})
}
