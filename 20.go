package main

/*
思路一： 用栈来实现
1. 左括号入栈
2. 右括号出栈
3. 遍历完成后栈为空，则有效否则无效

时间复杂度：O(n)
空间复杂度: O(n)
*/

func isValid(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)
	for i := 0; i < n; i++ {
		if v, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
