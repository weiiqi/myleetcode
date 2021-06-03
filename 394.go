package main

import (
	"strconv"
)

// 思路： 遍历S， 遇到"]" 取出字符，构建处对应的字符串，反转以后，压栈。其他字符直接入栈.遍历完，栈中就是所有的字符了。
func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ']':
			tmp := make([]byte, 0)
			for len(stack) != 0 && stack[len(stack)-1] != '[' {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				tmp = append(tmp, v)
			}

			stack = stack[:len(stack)-1]
			idx := 1
			for len(stack) >= idx && stack[len(stack)-idx] >= '0' && stack[len(stack)-idx] <= '9' {
				idx++
			}
			num := stack[len(stack)-idx+1:]
			stack = stack[:len(stack)-idx+1]
			count, _ := strconv.Atoi(string(num))
			for j := 0; j < count; j++ {
				for j := len(tmp) - 1; j >= 0; j-- {
					stack = append(stack, tmp[j])
				}
			}
		default:
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
