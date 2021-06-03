package main

import (
	"fmt"
)

/*
思路： 动态规划
state:(0,0) 到(x,y) 之间的距离为f(x, y)
function: f(x,y) = min(f(x+1, y), f(x,y+1))+nums(x,y)
initialize: f(0,0) = nums(0,0) f(i,0) = sum(0,0 -> i,0), f(0,i) = sum(0,0 -> 0,i)
answer: fun(n-1, m-1)
*/
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := 1; i < len(grid); i++ {
		grid[i][0] = grid[i][0] + grid[i-1][0]
	}
	for j := 1; j < len(grid[0]); j++ {
		grid[0][j] = grid[0][j] + grid[0][j-1]
	}
	tmp := make([]int, 0)
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[i]); j++ {
			min := min2(grid[i][j-1], grid[i-1][j])
			tmp = append(tmp, min)
			grid[i][j] = min + grid[i][j]
		}
	}
	fmt.Print(tmp)
	return grid[len(grid)-1][len(grid[0])-1]
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Print(minPathSum(nums))
}
