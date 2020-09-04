package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.octolab.org/pointer"

	. "github.com/kamilsk/lift/internal/model"
)

func TestCron_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Cron
		assert.NotPanics(t, func() { dst.Merge(Cron{Name: "cron"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Cron{Name: "cron-a"}
		assert.NotPanics(t, func() { dst.Merge(Cron{Name: "cron-b", Enabled: pointer.ToBool(true)}) })
		assert.Nil(t, dst.Enabled)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Cron{
			Name:     "importer",
			Enabled:  pointer.ToBool(false),
			Schedule: "11 17 * * *",
			Command:  "/bin/importer --count 10",
		}
		src := Cron{
			Name:     "importer",
			Enabled:  pointer.ToBool(true),
			Schedule: "11 * * * *",
			Command:  "/bin/importer --count 20",
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
		assert.Equal(t, Cron{
			Name:     "importer",
			Enabled:  pointer.ToBool(true),
			Schedule: "11 * * * *",
			Command:  "/bin/importer --count 20",
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

func TestCrons_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Crons
		assert.NotPanics(t, func() { dst.Merge(Crons{{Name: "cron"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Crons)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Crons{
			{
				Name:     "importer",
				Enabled:  pointer.ToBool(false),
				Schedule: "11 17 * * *",
				Command:  "/bin/importer --count 10",
			},
			{
				Name:     "cleaner",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 20 * * *",
				Command:  "/bin/cleaner --all",
			},
		}
		src := Crons{
			{
				Name:     "migrator",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 30 * * *",
				Command:  "/bin/migrator migrate",
			},
			{
				Name:     "importer",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 * * * *",
				Command:  "/bin/importer --count 20",
			},
		}

		dst.Merge(src)
		assert.Equal(t, Crons{
			{
				Name:     "importer",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 * * * *",
				Command:  "/bin/importer --count 20",
			},
			{
				Name:     "cleaner",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 20 * * *",
				Command:  "/bin/cleaner --all",
			},
			{
				Name:     "migrator",
				Enabled:  pointer.ToBool(true),
				Schedule: "11 30 * * *",
				Command:  "/bin/migrator migrate",
			},
		}, dst)
	})
}

func TestCrons_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Crons
		expected Crons
	}{
		"sorted": {
			input: Crons{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Crons{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Crons{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Crons{
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
