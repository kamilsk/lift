package model_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/kamilsk/lift/internal/model"
)

func TestQueue_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Queue
		assert.NotPanics(t, func() { dst.Merge(Queue{Name: "queue"}) })
		assert.Nil(t, dst)
	})

	t.Run("inappropriate source", func(t *testing.T) {
		var dst = Queue{Name: "queue-a"}
		assert.NotPanics(t, func() { dst.Merge(Queue{Name: "queue-b", DLQ: []string{"10m"}}) })
		assert.Nil(t, dst.DLQ)
	})

	t.Run("simple", func(t *testing.T) {
		dst := Queue{
			Name:    "queue",
			DLQ:     []string{"10m", "1h"},
			Aliases: []string{"queue.alias.1"},
		}
		src := Queue{
			Name:    "queue",
			DLQ:     []string{"20m"},
			Aliases: []string{"queue.alias.1", "queue.alias.2"},
		}

		dst.Merge(src)
		assert.Equal(t, Queue{
			Name:    "queue",
			DLQ:     []string{"10m", "1h", "20m"},
			Aliases: []string{"queue.alias.1", "queue.alias.2"},
		}, dst)
	})
}

func TestQueues_Merge(t *testing.T) {
	t.Run("nil destination", func(t *testing.T) {
		var dst *Queues
		assert.NotPanics(t, func() { dst.Merge(Queues{{Name: "queue"}}) })
		assert.Nil(t, dst)
	})

	t.Run("nil source", func(t *testing.T) {
		var dst = new(Queues)
		assert.NotPanics(t, func() { dst.Merge(nil) })
		assert.Empty(t, dst)
	})

	t.Run("with duplicates", func(t *testing.T) {
		dst := Queues{
			{
				Name: "queue-a",
				DLQ:  []string{"10m"},
			},
			{
				Name:    "queue-c",
				DLQ:     []string{"1m"},
				Aliases: []string{"alias.queue.1"},
			},
		}
		src := Queues{
			{
				Name: "queue-c",
				DLQ:  []string{"1m", "10m"},
			},
			{
				Name: "queue-b",
			},
		}

		dst.Merge(src)
		assert.Equal(t, Queues{
			{
				Name: "queue-a",
				DLQ:  []string{"10m"},
			},
			{
				Name:    "queue-c",
				DLQ:     []string{"1m", "10m"},
				Aliases: []string{"alias.queue.1"},
			},
			{
				Name: "queue-b",
			},
		}, dst)
	})
}

func TestQueues_Sort(t *testing.T) {
	tests := map[string]struct {
		input    Queues
		expected Queues
	}{
		"sorted": {
			input: Queues{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
			expected: Queues{
				{Name: "a"},
				{Name: "b"},
				{Name: "c"},
			},
		},
		"unsorted": {
			input: Queues{
				{Name: "b"},
				{Name: "c"},
				{Name: "a"},
			},
			expected: Queues{
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
