package linked_list

type Queue struct {
	inner *LinkedList
}

func NewQueue() *Queue {
	return &Queue{
		inner: NewLinkedList(),
	}
}

func (q *Queue) Empty() bool {
	return q.inner.Empty()
}

func (q *Queue) Enqueue(val int) {
	q.inner.PushBack(val)
}

func (q *Queue) Dequeue() int {
	if q.Empty() {
		panic("queue is empty")
	}

	return q.inner.PopFront()
}
