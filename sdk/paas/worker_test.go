package paas_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/sdk/paas"
)

func TestWorker_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Worker
		assert.NotPanics(t, func() { dst.Merge(Worker{Name: "worker"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Worker{Name: "worker-a"}
		assert.NotPanics(t, func() { dst.Merge(Worker{Name: "worker-b", Enabled: pointer.ToBool(true)}) })
		assert.Nil(t, dst.Enabled)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Worker{
			Name:          "worker",
			Enabled:       pointer.ToBool(false),
			Replicas:      1,
			LivenessProbe: "/usr/bin/worker check",
			Size:          "small",
		}
		src := Worker{
			Name:    "worker",
			Enabled: pointer.ToBool(true),
			Command: "/usr/bin/worker do",
			Commands: []string{
				"/usr/bin/worker do step",
				"/usr/bin/worker do action",
			},
			Replicas:      3,
			LivenessProbe: "/usr/bin/worker do check",
			Size:          "medium",
			Resources: &Resources{
				Requests: &Resource{
					CPU:    1,
					Memory: 10,
				},
				Limits: &Resource{
					CPU:    2,
					Memory: 20,
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Worker{
			Name:    "worker",
			Enabled: pointer.ToBool(true),
			Command: "/usr/bin/worker do",
			Commands: []string{
				"/usr/bin/worker do step",
				"/usr/bin/worker do action",
			},
			Replicas:      3,
			LivenessProbe: "/usr/bin/worker do check",
			Size:          "medium",
			Resources: &Resources{
				Requests: &Resource{
					CPU:    1,
					Memory: 10,
				},
				Limits: &Resource{
					CPU:    2,
					Memory: 20,
				},
			},
		}, dst)
	})
}

func TestWorkers_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Workers
		assert.NotPanics(t, func() { dst.Merge(Workers{{Name: "worker"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Workers)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicates", func(t *testing.T) {
		dst := Workers{
			{
				Name:          "worker-a",
				Enabled:       pointer.ToBool(false),
				Replicas:      1,
				LivenessProbe: "/usr/bin/worker check",
				Size:          "small",
			},
			{
				Name:          "worker-c",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/worker do",
				Replicas:      3,
				LivenessProbe: "/usr/bin/worker run",
				Size:          "small",
			},
		}
		src := Workers{
			{
				Name:          "worker-b",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/worker do",
				Replicas:      1,
				LivenessProbe: "/usr/bin/worker sleep",
				Size:          "small",
			},
			{
				Name:    "worker-a",
				Enabled: pointer.ToBool(true),
				Commands: []string{
					"/usr/bin/worker do step",
					"/usr/bin/worker do action",
				},
				Replicas:      3,
				LivenessProbe: "/usr/bin/worker do check",
				Size:          "medium",
				Resources: &Resources{
					Requests: &Resource{
						CPU:    1,
						Memory: 10,
					},
					Limits: &Resource{
						CPU:    2,
						Memory: 20,
					},
				},
			},
		}

		dst.Merge(src)
		assert.Equal(t, Workers{
			{
				Name:    "worker-a",
				Enabled: pointer.ToBool(true),
				Commands: []string{
					"/usr/bin/worker do step",
					"/usr/bin/worker do action",
				},
				Replicas:      3,
				LivenessProbe: "/usr/bin/worker do check",
				Size:          "medium",
				Resources: &Resources{
					Requests: &Resource{
						CPU:    1,
						Memory: 10,
					},
					Limits: &Resource{
						CPU:    2,
						Memory: 20,
					},
				},
			},
			{
				Name:          "worker-c",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/worker do",
				Replicas:      3,
				LivenessProbe: "/usr/bin/worker run",
				Size:          "small",
			},
			{
				Name:          "worker-b",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/worker do",
				Replicas:      1,
				LivenessProbe: "/usr/bin/worker sleep",
				Size:          "small",
			},
		}, dst)
	})
}

func TestWorkers_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Workers
		expected Workers
	}{
		"sorted": {
			input: Workers{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Workers{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Workers{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Workers{
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
