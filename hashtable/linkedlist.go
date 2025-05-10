package main

import "fmt"

type Node struct {
	Key  HtItem
	Next *Node
}

func NewNode(key HtItem) *Node {
	return &Node{key, nil}
}

type LinkedList struct {
	Head *Node
	Size int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

func (l *LinkedList) PushFront(key HtItem) {
	l.Size++

	if l.Head == nil {
		l.Head = NewNode(key)
		return
	}

	n := NewNode(key)
	n.Next = l.Head

	l.Head = n
}

func (l *LinkedList) SearchByKeyOfHtItem(key int) bool {
	if l.Head == nil {
		return false
	}

	p := l.Head
	for p != nil {
		if p.Key.Key == key {
			return true
		}

		p = p.Next
	}

	return false
}

func (l *LinkedList) DeleteByHtItemKey(key int) {
	if l.Head == nil {
		return
	}

	p := l.Head
	parent := l.Head

	for p != nil && p.Key.Key != key {
		parent = p
		p = p.Next
	}

	if p != nil {
		l.Size--
		parent.Next = p.Next
	}
}

func (l *LinkedList) Print() {
	p := l.Head
	for p != nil {
		fmt.Printf("%d:%d\t", p.Key.Key, p.Key.Value)
		p = p.Next
	}

	fmt.Println()
}
