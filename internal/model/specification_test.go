package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSpecification_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Specification
		assert.NotPanics(t, func() { dst.Merge(&Specification{Name: "service"}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Specification)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Specification{
			Name:     "base",
			Replicas: 1,
		}
		src := Specification{
			Name:        "service",
			Description: "Awesome service.",
			Kind:        "business",
			Host:        "http://example.com",
			Replicas:    3,

			Balancer: new(Balancer),
			Engine:   new(Engine),
			Logger:   new(Logger),
			SFTP:     new(SFTP),

			Elastic:      new(ElasticSearch),
			MongoDB:      new(MongoDB),
			PostgreSQL:   new(PostgreSQL),
			RabbitMQ:     new(RabbitMQ),
			Redis:        new(Redis),
			RedisSharded: new(ShardedRedis),
			Sphinxes:     make(Sphinxes, 3),

			Crons:        make(Crons, 3),
			Dependencies: make(Dependencies, 3),
			EnvVars:      EnvironmentVariables{"KEY": "value"},
			Executable:   make(Executable, 3),
			Proxies:      make(Proxies, 3),
			Queues:       make(Queues, 3),
			Workers:      make(Workers, 3),
		}

		dst.Merge(&src)
		assert.Equal(t, Specification{
			Name:        "service",
			Description: "Awesome service.",
			Kind:        "business",
			Host:        "http://example.com",
			Replicas:    3,

			Balancer: new(Balancer),
			Engine:   new(Engine),
			Logger:   new(Logger),
			SFTP:     new(SFTP),

			Elastic:      new(ElasticSearch),
			MongoDB:      new(MongoDB),
			PostgreSQL:   new(PostgreSQL),
			RabbitMQ:     new(RabbitMQ),
			Redis:        new(Redis),
			RedisSharded: new(ShardedRedis),
			Sphinxes:     make(Sphinxes, 1),

			Crons:        make(Crons, 1),
			Dependencies: make(Dependencies, 1),
			EnvVars:      EnvironmentVariables{"KEY": "value"},
			Executable:   make(Executable, 1),
			Proxies:      make(Proxies, 1),
			Queues:       make(Queues, 1),
			Workers:      make(Workers, 1),
		}, dst)
	})
}
