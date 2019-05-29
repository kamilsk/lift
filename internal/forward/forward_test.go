package forward_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/forward"
)

func TestExtractPort(t *testing.T) {
	tests := []struct {
		name       string
		definition string
		expected   uint16
	}{
		{"postgresql", "5432", 5432},
		{"mongodb", "mongodb://localhost:27017", 27017},
		{"rabbitmq", "localhost:5672", 5672},
		{"redis", "6379", 6379},
		{"sphinx", "9306", 9306},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			port, err := ExtractPort(tc.definition)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, port)
		})
	}
}

func TestPodName(t *testing.T) {
	tests := []struct {
		name            string
		service, entity string
		isLocal         bool
		expected        string
	}{
		{"local demo", "demo", "redis", true, "demo-local-redis-"},
		{"remote demo", "demo", "redis", false, "demo-redis-"},
	}
	for _, test := range tests {
		tc := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, tc.expected, PodName(tc.service, tc.entity, tc.isLocal))
		})
	}
}
