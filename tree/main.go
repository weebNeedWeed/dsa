package main

import "fmt"

func main() {
	fmt.Println("Creating a Binary Search Tree:")
	t := NewTree(nil)

	// Insert values to create a balanced BST
	values := []int{50, 30, 70, 20, 40, 60, 80, 15, 25, 35, 45, 55, 65, 75, 85}

	fmt.Println("Inserting values:", values)
	for _, val := range values {
		t.InsertRecursive(val)
	}

	fmt.Println("Inorder Traversal:")

	t.LNRStack()
}
