package main

import "fmt"

// 思路动态规划
// 满足以下条件可以使用动态规划
// 1: 求最大值
// 2: 求最小值
// 3: 求可行个数
// 满足不能排序或者交换
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 || len(triangle[0]) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	l := len(triangle)
	f := make([][]int, l)

	for i := 0; i < l; i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if f[i] == nil {
				f[i] = make([]int, len(triangle[i]))
			}
			f[i][j] = triangle[i][j]
		}
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			f[i][j] = min1(f[i+1][j], f[i+1][j+1]) + triangle[i][j]
		}
	}
	return f[0][0]
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	t := [][]int{
		{-10},
	}
	fmt.Print(minimumTotal(t))
}
