package linked_list

type ListNode struct {
	data int
	next *ListNode
}

type LinkedList struct {
	size int
	head *ListNode
	tail *ListNode
}

func NewNode(data int, next *ListNode) *ListNode {
	return &ListNode{
		data,
		next,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		size: 0,
		head: nil,
		tail: nil,
	}
}

// only push back and pop front are enough for our queue implementation
func (l *LinkedList) PushBack(val int) {
	l.size++
	if l.head == nil {
		n := NewNode(val, nil)
		l.head = n
		l.tail = n
		return
	}

	n := NewNode(val, nil)
	l.tail.next = n
	l.tail = n
}

func (l *LinkedList) Empty() bool {
	return l.head == nil
}

func (l *LinkedList) PopFront() int {
	if l.head == nil {
		panic("list is empty")
	}

	x := l.head.data
	l.head = l.head.next
	l.size--

	return x
}
