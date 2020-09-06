package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestExec_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Exec
		assert.NotPanics(t, func() { dst.Merge(Exec{Name: "binary"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Exec{Name: "binary-a"}
		assert.NotPanics(t, func() { dst.Merge(Exec{Name: "binary-b", Enabled: pointer.ToBool(true)}) })
		assert.Nil(t, dst.Enabled)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Exec{
			Name:          "binary",
			Enabled:       pointer.ToBool(false),
			Replicas:      1,
			LivenessProbe: "/usr/bin/binary run",
			Size:          "small",
		}
		src := Exec{
			Name:          "binary",
			Enabled:       pointer.ToBool(true),
			Command:       "./entrypoint.sh -m http.server 8890",
			Image:         "python",
			Replicas:      3,
			Port:          8890,
			RedinessProbe: "curl --fail http://127.0.0.1:8890",
			LivenessProbe: "curl --fail http://127.0.0.1:8890",
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
		assert.Equal(t, Exec{
			Name:          "binary",
			Enabled:       pointer.ToBool(true),
			Command:       "./entrypoint.sh -m http.server 8890",
			Image:         "python",
			Replicas:      3,
			Port:          8890,
			RedinessProbe: "curl --fail http://127.0.0.1:8890",
			LivenessProbe: "curl --fail http://127.0.0.1:8890",
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

func TestExecutable_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Executable
		assert.NotPanics(t, func() { dst.Merge(Executable{{Name: "binary"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Executable)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicates", func(t *testing.T) {
		dst := Executable{
			{
				Name:          "binary-a",
				Enabled:       pointer.ToBool(false),
				Replicas:      1,
				LivenessProbe: "/usr/bin/binary run",
				Size:          "small",
			},
			{
				Name:          "binary-c",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/binary do",
				Replicas:      3,
				LivenessProbe: "/usr/bin/binary run",
				Size:          "small",
			},
		}
		src := Executable{
			{
				Name:          "binary-b",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/binary do",
				Replicas:      1,
				LivenessProbe: "/usr/bin/binary sleep",
				Size:          "small",
			},
			{
				Name:          "binary-a",
				Enabled:       pointer.ToBool(true),
				Command:       "./entrypoint.sh -m http.server 8890",
				Image:         "python",
				Replicas:      3,
				Port:          8890,
				RedinessProbe: "curl --fail http://127.0.0.1:8890",
				LivenessProbe: "curl --fail http://127.0.0.1:8890",
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
		assert.Equal(t, Executable{
			{
				Name:          "binary-a",
				Enabled:       pointer.ToBool(true),
				Command:       "./entrypoint.sh -m http.server 8890",
				Image:         "python",
				Replicas:      3,
				Port:          8890,
				RedinessProbe: "curl --fail http://127.0.0.1:8890",
				LivenessProbe: "curl --fail http://127.0.0.1:8890",
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
				Name:          "binary-c",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/binary do",
				Replicas:      3,
				LivenessProbe: "/usr/bin/binary run",
				Size:          "small",
			},
			{
				Name:          "binary-b",
				Enabled:       pointer.ToBool(true),
				Command:       "/usr/bin/binary do",
				Replicas:      1,
				LivenessProbe: "/usr/bin/binary sleep",
				Size:          "small",
			},
		}, dst)
	})
}

func TestExecutable_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Executable
		expected Executable
	}{
		"sorted": {
			input: Executable{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Executable{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Executable{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Executable{
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
