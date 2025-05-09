package main

import (
	"strconv"
	"strings"
)

func isOperator(elm string) (bool, string) {
	switch elm {
	case "+", "-", "*", "/":
		return true, elm
	}
	return false, ""
}

func evalExp(exp string) int {
	s := newStack(100)
	elms := strings.Fields(exp)
	var temp int
	for _, elm := range elms {
		if isOp, op := isOperator(elm); isOp {
			a := s.pop()
			b := s.pop()
			if op == "+" {
				temp = a + b
			}
			if op == "-" {
				temp = a - b
			}
			if op == "*" {
				temp = a * b
			}
			if op == "/" {
				temp = b / a
			}
			s.push(temp)
		} else {
			num, _ := strconv.Atoi(elm)
			s.push(num)
		}
	}
	temp = s.pop() // result
	return temp
}
