package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestRabbitMQ_Merge(t *testing.T) {
	t.Run("nil rabbitmq", func(t *testing.T) {
		var rabbitmq *RabbitMQ
		assert.NotPanics(t, func() { rabbitmq.Merge(&RabbitMQ{Version: "3.6"}) })
		assert.Nil(t, rabbitmq)
	})
}
