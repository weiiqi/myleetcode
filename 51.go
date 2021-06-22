package main

/*
思路一: 回溯算法
时间复杂度: O (n!) 第一行有N个选择，第二行有n-1个选择，一次类推，为N！
空间复杂度: O(4n)
	1。 递归深度最大为N
	2。 存放列数据， 存放左对角线，存放右对角线，最大都为N空间
	3。 和为4N，简化为O(n)
*/
func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)
	queens := make([]int, 0)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	colums := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backTrackNQ(queens, n, 0, colums, diagonals1, diagonals2, &ret)
	return ret
}

func backTrackNQ(queen []int, n, row int, colums, diagonals1, diagonals2 map[int]bool, ret *[][]string) {
	// drill down
	if row == n {
		board := generateBoard(queen, n)
		*ret = append((*ret), board)
		return
	}
	// process current level logic
	for i := 0; i < n; i++ {
		if colums[i] { // 跳过当前列已经存在的数据
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] { // 跳过对角线已存在的数据. 左右对角线，要么和相等，要么差相等
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] { // 跳过对角线已存在的数据. 左右对角线，要么和相等，要么差相等
			continue
		}
		queen[row] = i
		colums[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		// drill down
		backTrackNQ(queen, n, row+1, colums, diagonals1, diagonals2, ret)

		// reverse state
		queen[row] = -1
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
		delete(colums, i)
	}
}

// 棋盘生成，只在queen的位置替换为皇后即可
func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}
