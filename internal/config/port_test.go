package config_test

import (
	"testing"

	. "github.com/kamilsk/lift/internal/config"
	"github.com/stretchr/testify/assert"
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
