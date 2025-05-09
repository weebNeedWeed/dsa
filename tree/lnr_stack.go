package main

import "fmt"

func (t *Tree) LNRStack() {
	if t.Root == nil {
		return
	}

	s := make([]*Node, 0)

	p := t.Root
	for {
		for p != nil {
			s = append(s, p)
			p = p.Left
		}

		if len(s) != 0 {
			pop := s[len(s)-1]
			s = s[:len(s)-1]

			fmt.Println(pop.Key)

			if pop.Right != nil {
				s = append(s, pop.Right)
			}
		} else {
			break
		}
	}
}
