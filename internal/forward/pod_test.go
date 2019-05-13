package forward_test

import (
	"testing"

	. "github.com/kamilsk/lift/internal/forward"
	"github.com/stretchr/testify/assert"
)

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
