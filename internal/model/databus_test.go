package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestDataBus_Merge(t *testing.T) {
	t.Run("nil databus destination", func(t *testing.T) {
		var databus *DataBus
		assert.NotPanics(t, func() { databus.Merge(&DataBus{BatchSize: 10}) })
		assert.Nil(t, databus)
	})

	t.Run("nil databus source", func(t *testing.T) {
		var databus = new(DataBus)
		assert.NotPanics(t, func() { databus.Merge(nil) })
		assert.Empty(t, databus)
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
			Schemas:   []string{"schema-a", "schema-b", "schema-c", "schema-d"},
		}, dst)
	})
}
