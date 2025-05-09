package circular_buffer

type Queue struct {
	inner    []int
	read     int
	write    int
	capacity int
}

func NewQueue(size int) Queue {
	return Queue{
		inner:    make([]int, size),
		read:     0,
		write:    0,
		capacity: size,
	}
}

func (q *Queue) Empty() bool {
	return q.read == q.write
}

func (q *Queue) Full() bool {
	return q.getNextPointerValue(q.write) == q.read
}

func (q *Queue) getNextPointerValue(p int) int {
	return (p + 1) % q.capacity
}

func (q *Queue) Enqueue(val int) {
	if q.Full() {
		panic("queue is full")
	}

	q.inner[q.write] = val
	q.write = q.getNextPointerValue(q.write)
}

func (q *Queue) Dequeue() int {
	if q.Empty() {
		panic("queue is empty")
	}

	x := q.inner[q.read]
	q.read = q.getNextPointerValue(q.read)
	return x
}
