package main

import "fmt"

// 思路： dfs 深度优先遍历
func numIslands(grid [][]byte) int {
	var count int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && dfs(grid, i, j) >= 1 {
				count++
			}
		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return 0
	}
	if grid[i][j] == '1' {
		grid[i][j] = 0
		return dfs(grid, i-1, j) + dfs(grid, i, j-1) + dfs(grid, i+1, j) + dfs(grid, i, j+1) + 1
	}
	return 0
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	//fmt.Println(numIslands(grid))
	fmt.Println(numIslands_20210630(grid))
}

/*
二刷code-200, 2021-06-30
思路一：DFS
本题目是基于网格的深度优先遍历
*/

/*
模版：
*/
func templateDFSGrid(grid [][]byte, r, c int) {
	// terminate
	if !inArea(grid, r, c) {
		return
	}
	// process current level logic
	// 如果不是1， 即为非陆地
	if grid[r][c] != '1' {
		return
	}
	// 2 代表访问过的陆地节点， 避免重复访问
	grid[r][c] = '2'
	// drill down
	templateDFSGrid(grid, r-1, c)
	templateDFSGrid(grid, r+1, c)
	templateDFSGrid(grid, r, c-1)
	templateDFSGrid(grid, r, c+1)
}

func inArea(grid [][]byte, r, c int) bool {
	return 0 <= r && r < len(grid) && c >= 0 && c < len(grid[0])
}

/*
时间复杂度：O(n) -> 格子数量，每个格子只会被遍历一次
空间复杂度: O(n) -> 格子数量，最大递归深度为所有节点都在递归中
*/
func numIslands_20210630(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				templateDFSGrid(grid, i, j)
				count++
			}
		}
	}
	return count
}
