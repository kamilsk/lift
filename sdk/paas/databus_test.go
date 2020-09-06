package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestDataBus_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *DataBus
		assert.NotPanics(t, func() { dst.Merge(&DataBus{BatchSize: 10}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(DataBus)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicate schemas", func(t *testing.T) {
		dst := DataBus{
			BatchSize: 1,
			Schemas:   []string{"schema-c", "schema-a", "schema-b"},
		}
		src := DataBus{
			BatchSize: 10,
			Schemas:   []string{"schema-b", "schema-d"},
		}

		dst.Merge(&src)
		assert.Equal(t, DataBus{
			BatchSize: 10,
			Schemas:   []string{"schema-c", "schema-a", "schema-b", "schema-d"},
		}, dst)
	})
}
