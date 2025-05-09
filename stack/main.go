package main

import "fmt"

type stack struct {
	innerArr []int
	top      int
}

func newStack(size int) stack {
	return stack{
		innerArr: make([]int, size),
		top:      -1,
	}
}

func (s *stack) push(elm int) {
	s.top++
	s.innerArr[s.top] = elm
}

func (s *stack) pop() int {
	if s.top < 0 {
		panic("stack is empty")
	}

	r := s.innerArr[s.top]
	s.top--

	return r
}

func (s *stack) empty() bool {
	return s.top == -1
}

func convertDecToBin(x int) {
	s := newStack(100)

	for x != 0 {
		s.push(x % 2)
		x /= 2
	}

	for !s.empty() {
		fmt.Print(s.pop())
	}
}

func main() {
	exp := "10 3 + 2 * 13 /"

	fmt.Println(evalExp(exp))
}
