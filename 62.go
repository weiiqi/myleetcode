package main

import "fmt"

/*
思路： 动态规划
state：f(x,y) 表示从起点到(x,y) 之间的不同路径条目数
function: f(x,y) = f(x-1,y) + f(x, y-1)
init: f(x,i) -> f(0,0) .. f(0, i) = 1, f(i,y) = f(0,0) .. f(i,0) =1
answer: f(m-1,n-1)
*/
func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Print(uniquePaths(3, 2))
}
