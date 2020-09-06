package paas_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestHost_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Host
		assert.NotPanics(t, func() { dst.Merge(Host{Name: "host"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Host{Name: "host-a"}
		assert.NotPanics(t, func() { dst.Merge(Host{Name: "host-b", AgentPort: 1234}) })
		assert.Empty(t, dst.AgentPort)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Host{
			Name:        "host",
			AgentPort:   1234,
			Connections: 1,
			MaxConns:    2,
			Weight:      3,
			Backup:      pointer.ToBool(false),
		}
		src := Host{
			Name:        "host",
			AgentPort:   4321,
			Connections: 10,
			MaxConns:    20,
			Weight:      30,
			Backup:      pointer.ToBool(true),
		}

		dst.Merge(src)
		assert.Equal(t, Host{
			Name:        "host",
			AgentPort:   4321,
			Connections: 10,
			MaxConns:    20,
			Weight:      30,
			Backup:      pointer.ToBool(true),
		}, dst)
	})
}

func TestHosts_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Hosts
		assert.NotPanics(t, func() { dst.Merge(Hosts{{Name: "host-a"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Hosts)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicates", func(t *testing.T) {
		dst := Hosts{
			{
				Name:        "host-a",
				AgentPort:   1234,
				Connections: 1,
				MaxConns:    2,
				Weight:      3,
			},
			{
				Name:        "host-b",
				AgentPort:   1234,
				Connections: 1,
				MaxConns:    2,
				Weight:      3,
				Backup:      pointer.ToBool(true),
			},
		}
		src := Hosts{
			{
				Name:        "host-b",
				AgentPort:   4321,
				Connections: 10,
				MaxConns:    20,
				Weight:      30,
				Backup:      pointer.ToBool(false),
			},
			{
				Name:        "host-c",
				AgentPort:   1234,
				Connections: 1,
				MaxConns:    2,
				Weight:      3,
			},
		}

		dst.Merge(src)
		assert.Equal(t, Hosts{
			{
				Name:        "host-a",
				AgentPort:   1234,
				Connections: 1,
				MaxConns:    2,
				Weight:      3,
			},
			{
				Name:        "host-b",
				AgentPort:   4321,
				Connections: 10,
				MaxConns:    20,
				Weight:      30,
				Backup:      pointer.ToBool(false),
			},
			{
				Name:        "host-c",
				AgentPort:   1234,
				Connections: 1,
				MaxConns:    2,
				Weight:      3,
			},
		}, dst)
	})
}

func TestHosts_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Hosts
		expected Hosts
	}{
		"sorted": {
			input: Hosts{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Hosts{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Hosts{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Hosts{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			sort.Sort(test.input)
			assert.Equal(t, test.expected, test.input)
		})
	}
}
