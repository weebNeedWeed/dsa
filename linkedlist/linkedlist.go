package main

import "log"

type linkedList struct {
	head *node
	size int
}

func newLinkedList(head *node) *linkedList {
	return &linkedList{head, 0}
}

func (l *linkedList) getSize() int {
	return l.size
}

func (l *linkedList) empty() bool {
	return l.size == 0
}

func (l *linkedList) pushFront(data int) {
	n := newNode(data, l.head)
	l.head = n
	l.size++
}

func (l *linkedList) valueAt(index int) int {
	if index < 0 {
		log.Panicf("index is out of bound, %d", index)
	}

	// because the next for statement cannot be merged with this case
	if index == 0 && l.head != nil {
		return l.head.data
	}

	i := 0
	p := l.head
	for p != nil && i < index {
		i++
		p = p.next
	}

	// the element for given index was not found
	if i != index || p == nil {
		log.Panicf("index if out of bound, %d", index)
	}

	return p.data
}

func (l *linkedList) pushBack(data int) {
	n := newNode(data, nil)

	if l.head == nil { // if list is nil, we directly set the head
		l.head = n
	} else {
		last := l.head
		for last.next != nil {
			last = last.next
		}

		last.next = n
	}

	l.size++
}

func (l *linkedList) insert(index int, value int) {
	if index < 0 || index >= l.size {
		log.Panicf("index is out of bound, %d", index)
	}

	if index == 0 {
		l.pushFront(value)
		return
	}

	// we need to find the index-1 node
	p := l.head
	i := 0
	for p != nil && i < index-1 {
		p = p.next
		i++
	}

	// rarely happend
	// if p == nil {
	// 	log.Panicf("index is out of bound, %d", index)
	// }

	// now p is at index - 1, n is at index
	n := newNode(value, p.next)
	p.next = n
	l.size++
}

func (l *linkedList) popFirst() int {
	if l.head == nil {
		log.Panic("head is nil")
	}

	res := l.head.data
	l.head = l.head.next
	l.size--
	return res
}

func (l *linkedList) popBack() int {
	if l.head == nil {
		log.Panic("head is nil")
	}

	l.size--
	if l.head.next == nil { // if list has only 1 items, the prev pointer could not work
		res := l.head.data
		l.head = nil
		return res
	}

	var prev *node = nil // save the node placed before last node
	last := l.head
	for last.next != nil {
		prev = last
		last = last.next
	}

	res := last.data
	prev.next = last.next

	return res
}

// for testing purposes
func (l *linkedList) toArray() []int {
	result := []int{}

	p := l.head
	for p != nil {
		result = append(result, p.data)
		p = p.next
	}

	return result
}

func main() {
}
