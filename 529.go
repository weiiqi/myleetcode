package main

import "fmt"

/*
思路一： 深度优先遍历
每次扩展八个方向，遇到存在地雷的边就停下，否者递归下去
时间复杂度: O(n) 棋盘每个都遍历一次
空间复杂度: O(n) 最坏的情况递归深度为棋盘格子数
*/

func main() {
	board := [][]byte{
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'M', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
		{'E', 'E', 'E', 'E', 'E'},
	}
	fmt.Println(updateBoard(board, []int{3, 0}))
}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs_529(x, y, board)
	}
	return board
}

var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

func dfs_529(x, y int, board [][]byte) {
	// terminate
	// 这里的终止条件，需要判断是否==E，如果==E，则说明是已经访问过的节点，否者会死循环
	if !inBoard(board, x, y) || board[x][y] != 'E' {
		return
	}
	// process current level logic
	cnt := 0
	for i := 0; i < 8; i++ {
		tx, ty := x+dirX[i], y+dirY[i]
		if !inBoard(board, tx, ty) {
			continue
		}
		if board[tx][ty] == 'M' {
			cnt++
		}
	}
	if cnt > 0 {
		board[x][y] = byte(cnt + '0')
	} else {
		board[x][y] = 'B'
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			dfs_529(tx, ty, board)
		}
	}
}

func inBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}
