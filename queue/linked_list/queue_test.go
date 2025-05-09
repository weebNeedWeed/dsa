package linked_list

import "testing"

func TestQueue_EnqueueDequeue(t *testing.T) {
	q := NewQueue()

	// Test that a new queue is empty
	if !q.Empty() {
		t.Error("Expected new queue to be empty")
	}

	// Enqueue elements and test non-empty state
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		q.Enqueue(v)
	}
	if q.Empty() {
		t.Error("Expected queue not to be empty after enqueues")
	}

	// Dequeue and verify FIFO order
	for i, expected := range values {
		got := q.Dequeue()
		if got != expected {
			t.Errorf("Dequeue #%d: expected %d, got %d", i+1, expected, got)
		}
	}

	// After every element is dequeued, queue should be empty
	if !q.Empty() {
		t.Error("Expected queue to be empty after dequeues")
	}
}

func TestQueue_DequeueEmptyPanics(t *testing.T) {
	q := NewQueue()

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when dequeuing from an empty queue")
		}
	}()
	// This should panic because the queue is empty
	q.Dequeue()
}

func TestQueue_MixedOperations(t *testing.T) {
	q := NewQueue()

	// Enqueue a few items and dequeue one
	q.Enqueue(10)
	q.Enqueue(20)
	result := q.Dequeue()
	if result != 10 {
		t.Errorf("Expected first dequeue to be 10, got %d", result)
	}

	// Enqueue another item and continue dequeueing
	q.Enqueue(30)
	result = q.Dequeue()
	if result != 20 {
		t.Errorf("Expected second dequeue to be 20, got %d", result)
	}
	result = q.Dequeue()
	if result != 30 {
		t.Errorf("Expected third dequeue to be 30, got %d", result)
	}
}
