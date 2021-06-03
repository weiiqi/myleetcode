package main

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	stack := make([]int, 0)
	for _, v := range tokens {
		switch v {
		case "+", "-", "*", "/":
			if len(stack) < 2 {
				return -1
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var result int
			switch v {
			case "*":
				result = a * b
			case "+":
				result = a + b
			case "/":
				result = a / b
			case "-":
				result = a - b
			}
			stack = append(stack, result)
		default:
			val, _ := strconv.Atoi(v)
			stack = append(stack, val)
		}
	}
	return stack[0]
}
