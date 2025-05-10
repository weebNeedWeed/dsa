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

func lnr(n *Node) {
	if n == nil {
		return
	}

	lnr(n.Left)
	fmt.Println(n.Key)
	lnr(n.Left)
}

func (t *Tree) Lnr() {
	if t.Root == nil {
		return
	}

	lnr(t.Root)
}

func countNodes(n *Node, noChild, oneChild, twoChilds *int) {
	if n == nil {
		return
	}

	countNodes(n.Left, noChild, oneChild, twoChilds)

	if n.Left == nil && n.Right == nil {
		*noChild = *noChild + 1
	} else if n.Left == nil || n.Right == nil {
		*oneChild++
	} else {
		*twoChilds++
	}

	countNodes(n.Right, noChild, oneChild, twoChilds)
}

func (t *Tree) CountNodes() (int, int, int) {
	no, one, two := 0, 0, 0

	countNodes(t.Root, &no, &one, &two)

	return no, one, two
}

func (t *Tree) Insert(key int) {
	if t.Root == nil {
		t.Root = NewNode(key)
		return
	}

	insertNode(t.Root, key)
}

func insertNode(n *Node, key int) *Node {
	if n == nil {
		return NewNode(key)
	}

	if n.Key == key {
		return n
	}

	if key > n.Key {
		n.Right = insertNode(n.Right, key)
	} else {
		n.Left = insertNode(n.Left, key)
	}

	return n
}

func calHeight(n *Node) int {
	if n == nil {
		return -1
	}

	left := calHeight(n.Left) + 1
	right := calHeight(n.Right) + 1

	return max(left, right)
}

func (t *Tree) CalHeight() int {
	return calHeight(t.Root)
}
