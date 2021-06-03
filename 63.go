package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i] == nil {
				dp[i] = make([]int, n)
			}
			dp[i][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 || dp[i-1][0] == 0 {
			dp[i][0] = 0
		}
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 || dp[0][i-1] == 0 {
			dp[0][i] = 0
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func main() {
	nums := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	fmt.Print(uniquePathsWithObstacles(nums))
}
