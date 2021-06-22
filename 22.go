package main

import "fmt"

/*
思路一: 尝试递归
时间复杂度: O(n^2)
空间复杂度: O(n)
*/

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	generate(0, 0, n, "", &ret)
	return ret
}

func generate(left, right, max int, s string, ret *[]string) {
	// recursion terminator
	if left == max && right == max {
		*ret = append(*ret, s)
		return
	}
	// process logins in current level
	// drill down
	if left < max {
		s1 := s + "("
		generate(left+1, right, max, s1, ret)
	}
	if right < left {
		s2 := s + ")"
		generate(left, right+1, max, s2, ret)
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
