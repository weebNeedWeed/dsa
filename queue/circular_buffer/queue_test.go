package circular_buffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	q := NewQueue(3)
	t.Run("should enqueue successfully", func(t *testing.T) {
		q.Enqueue(1)
		q.Enqueue(2)

		if q.inner[0] != 1 || q.inner[1] != 2 || q.read != 0 || q.write != 2 {
			t.Errorf("failed to enqueue, got %v", q)
		}
	})

	t.Run("should panic when being full", func(t *testing.T) {
		assert.Panics(t, func() {
			q.Enqueue(3)
		}, "no panic when being full")
	})
}

func TestDequeue(t *testing.T) {
	q := NewQueue(2)
	q.Enqueue(1)
	t.Run("should dequeue successfully", func(t *testing.T) {
		got := q.Dequeue()
		want := 1

		if got != want {
			t.Errorf("failed to dequeue; want %v, got %v", want, got)
		}
	})

	t.Run("should panic when being empty", func(t *testing.T) {
		assert.Panics(t, func() {
			q.Dequeue()
		}, "no panic when being empty")
	})
}
