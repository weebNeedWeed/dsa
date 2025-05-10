package main

import "fmt"

func main() {
	t := NewTree(nil)

	// Create root node
	t.Root = NewNode(50)

	// Insert more nodes to create a binary search tree
	t.Insert(30)
	t.Insert(70)
	t.Insert(20)
	t.Insert(40)
	t.Insert(60)
	t.Insert(80)
	t.Insert(90)

	fmt.Println("Binary Search Tree created with nodes: 50, 30, 70, 20, 40, 60, 80")

	// Count different types of nodes
	noChild, oneChild, twoChilds := t.CountNodes()
	fmt.Printf("Nodes with no children: %d\n", noChild)
	fmt.Printf("Nodes with one child: %d\n", oneChild)
	fmt.Printf("Nodes with two children: %d\n", twoChilds)

	fmt.Println("Height of the tree", t.CalHeight())
}
