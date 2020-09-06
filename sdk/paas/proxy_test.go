package paas_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestProxy_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Proxy
		assert.NotPanics(t, func() { dst.Merge(Proxy{Name: "proxy"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Proxy{Name: "proxy-a"}
		assert.NotPanics(t, func() { dst.Merge(Proxy{Name: "proxy-b", Enabled: pointer.ToBool(true)}) })
		assert.Nil(t, dst.Enabled)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Proxy{
			Name:    "proxy",
			Enabled: pointer.ToBool(false),
			Hosts: Hosts{
				{
					Name:      "host-a",
					AgentPort: 1234,
					MaxConns:  1,
					Backup:    pointer.ToBool(true),
				},
			},
		}
		src := Proxy{
			Name:    "proxy",
			Enabled: pointer.ToBool(true),
			Hosts: Hosts{
				{
					Name:      "host-a",
					AgentPort: 4321,
					Backup:    pointer.ToBool(false),
				},
				{
					Name:      "host-b",
					AgentPort: 1234,
					MaxConns:  1,
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Proxy{
			Name:    "proxy",
			Enabled: pointer.ToBool(true),
			Hosts: Hosts{
				{
					Name:      "host-a",
					AgentPort: 4321,
					MaxConns:  1,
					Backup:    pointer.ToBool(false),
				},
				{
					Name:      "host-b",
					AgentPort: 1234,
					MaxConns:  1,
				},
			},
		}, dst)
	})
}

func TestProxies_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Proxies
		assert.NotPanics(t, func() { dst.Merge(Proxies{{Name: "proxy"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Proxies)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Proxies{
			{
				Enabled: pointer.ToBool(false),
				Hosts: Hosts{
					{
						Name:      "host-a",
						AgentPort: 1234,
						MaxConns:  1,
						Backup:    pointer.ToBool(false),
					},
				},
			},
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-b",
						AgentPort: 1234,
						MaxConns:  1,
					},
				},
				Name: "proxy-a",
			},
		}
		src := Proxies{
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-c",
						AgentPort: 1234,
						MaxConns:  1,
					},
				},
				Name: "proxy-a",
			},
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-a",
						AgentPort: 4321,
						MaxConns:  10,
						Backup:    pointer.ToBool(true),
					},
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Proxies{
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-a",
						AgentPort: 4321,
						MaxConns:  10,
						Backup:    pointer.ToBool(true),
					},
				},
			},
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-b",
						AgentPort: 1234,
						MaxConns:  1,
					},
					{
						Name:      "host-c",
						AgentPort: 1234,
						MaxConns:  1,
					},
				},
				Name: "proxy-a",
			},
		}, dst)
	})
}

func TestProxies_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Proxies
		expected Proxies
	}{
		"sorted": {
			input: Proxies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Proxies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Proxies{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Proxies{
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
