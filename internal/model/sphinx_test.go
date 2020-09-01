package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestSphinxes_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Sphinxes
		assert.NotPanics(t, func() { dst.Merge(Sphinxes{{Name: "test"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Sphinxes)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Sphinxes{
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
				Haproxy: "2.0.1",
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
				Name: "sphinx-a",
			},
		}
		src := Sphinxes{
			{
				Enabled: pointer.ToBool(true),
				Hosts: Hosts{
					{
						Name:      "host-c",
						AgentPort: 1234,
						MaxConns:  1,
					},
				},
				Name: "sphinx-a",
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
				Haproxy: "2.0.14",
			},
		}

		dst.Merge(src)
		assert.Equal(t, Sphinxes{
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
				Haproxy: "2.0.14",
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
				Name: "sphinx-a",
			},
		}, dst)
	})
}

func TestSphinxes_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Sphinxes
		expected Sphinxes
	}{
		"sorted": {
			input: Sphinxes{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Sphinxes{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Sphinxes{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Sphinxes{
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
