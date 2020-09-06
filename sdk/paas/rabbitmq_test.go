package paas_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestRabbitMQ_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *RabbitMQ
		assert.NotPanics(t, func() { dst.Merge(&RabbitMQ{Version: "3.6"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(RabbitMQ)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := RabbitMQ{
			Version: "3.5",
			Size:    "small",
			Vhosts:  "vhost-b,vhost-d",
		}
		src := RabbitMQ{
			Enabled: pointer.ToBool(true),
			Version: "3.6",
			Size:    "medium",
			Vhosts:  "vhost-c,vhost-a,vhost-b",
		}

		dst.Merge(&src)
		assert.Equal(t, RabbitMQ{
			Enabled: pointer.ToBool(true),
			Version: "3.6",
			Size:    "medium",
			Vhosts:  "vhost-b,vhost-d,vhost-c,vhost-a",
		}, dst)
	})
}
