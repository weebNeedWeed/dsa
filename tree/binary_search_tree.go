package main

import "fmt"

type Node struct {
	Key         int
	Left, Right *Node
}

func NewNode(key int) *Node {
	return &Node{
		key, nil, nil,
	}
}

type Tree struct {
	Root *Node
}

func NewTree(root *Node) *Tree {
	return &Tree{
		root,
	}
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = NewNode(key)
		return
	}

	p := t.Root
	var parent *Node

	for p != nil && p.Key != key {
		parent = p
		if key > p.Key {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	if p == nil {
		if key > parent.Key {
			parent.Right = NewNode(key)
		} else {
			parent.Left = NewNode(key)
		}
	}
}

func (t *Tree) InsertRecursive(key int) {
	if t.Root == nil {
		t.Root = NewNode(key)
	} else {
		insert(t.Root, key)
	}
}

func insert(n *Node, key int) *Node {
	if n == nil {
		return NewNode(key)
	}

	if n.Key == key {
		return n
	}

	if key > n.Key {
		n.Right = insert(n.Right, key)
	} else {
		n.Left = insert(n.Left, key)
	}

	return n
}

func (t Tree) LNR() {
	lnr(t.Root)
}

func lnr(n *Node) {
	if n == nil {
		return
	}

	lnr(n.Left)
	fmt.Println(n.Key)
	lnr(n.Right)
}

func (t *Tree) Delete(key int) {
	if t.Root == nil {
		return
	}

	p := t.Root
	parent := t.Root

	for p != nil && p.Key != key {
		parent = p
		if key > p.Key {
			p = p.Right
		} else {
			p = p.Left
		}
	}

	if p == nil {
		return
	}

	if p.Left == nil {
		if p.Key > parent.Key {
			parent.Right = p.Right
		} else {
			parent.Left = p.Right
		}
	} else if p.Right == nil {
		if p.Key > parent.Key {
			parent.Right = p.Left
		} else {
			parent.Left = p.Left
		}
	} else {
		leftMost := getSuccessor(p.Right)
		t.Delete(leftMost.Key)
		p.Key = leftMost.Key
	}
}

func getSuccessor(n *Node) *Node {
	curr := n
	if curr != nil && curr.Left != nil {
		curr = curr.Left
	}

	return curr
}
