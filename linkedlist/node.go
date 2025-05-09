package main

type node struct {
	data int
	next *node
}

func newNode(data int, next *node) *node {
	return &node{
		data: data,
		next: next,
	}
}
