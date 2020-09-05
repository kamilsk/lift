package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestDependency_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Dependency
		assert.NotPanics(t, func() { dst.Merge(Dependency{Name: "service"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Dependency{Name: "service-a"}
		assert.NotPanics(t, func() { dst.Merge(Dependency{Name: "service-b", Mock: pointer.ToBool(true)}) })
		assert.Nil(t, dst.Mock)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Dependency{
			Name: "service",
		}
		src := Dependency{
			Name:         "service",
			Mock:         pointer.ToBool(true),
			MockReplicas: 3,
		}

		dst.Merge(src)
		assert.Equal(t, Dependency{
			Name:         "service",
			Mock:         pointer.ToBool(true),
			MockReplicas: 3,
		}, dst)
	})
}

func TestDependencies_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Dependencies
		assert.NotPanics(t, func() { dst.Merge(Dependencies{{Name: "service"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Dependencies)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Dependencies{
			{
				Name: "service-a",
				Mock: pointer.ToBool(true),
			},
			{
				Name:         "service-b",
				MockReplicas: 10,
			},
		}
		src := Dependencies{
			{
				Name: "service-a",
				Mock: pointer.ToBool(false),
			},
			{
				Name: "service-c",
			},
		}

		dst.Merge(src)
		assert.Equal(t, Dependencies{
			{
				Name: "service-a",
				Mock: pointer.ToBool(false),
			},
			{
				Name:         "service-b",
				MockReplicas: 10,
			},
			{
				Name: "service-c",
			},
		}, dst)
	})
}

func TestDependencies_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Dependencies
		expected Dependencies
	}{
		"sorted": {
			input: Dependencies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Dependencies{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Dependencies{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Dependencies{
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
