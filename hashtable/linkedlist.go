package main

type Node[T any] struct {
	Key  T
	Next *Node[T]
}

func NewNode[T any](key T) *Node[T] {
	return &Node[T]{key, nil}
}

type LinkedList[T any] struct {
	Head *Node[T]
	Size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{nil, 0}
}

// func (l *LinkedList) PushFront(key int) {
// 	l.Size++

// 	if l.Head == nil {
// 		l.Head = NewNode(key)
// 		return
// 	}

// 	n := NewNode(key)
// 	n.Next = l.Head.Next

// 	l.Head = n
// }

// func (l *LinkedList) Search(key int) bool {
// 	if l.Head == nil {
// 		return false
// 	}

// 	p := l.Head
// 	for p != nil {
// 		if p.Key == key {
// 			return true
// 		}

// 		p = p.Next
// 	}

// 	return false
// }
