package main

// 思路： 将二维数组转化为一维数组，进行二分搜索
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := len(matrix)
	col := len(matrix[0])
	start := 0
	end := row*col - 1
	for start+1 < end {
		mid := start + (end-start)/2
		val := matrix[mid/col][mid%col]
		if val > target {
			end = mid
		} else if val < target {
			start = mid
		} else {
			return true
		}
	}
	if matrix[start/col][start%col] == target || matrix[end/col][end%col] == target {
		return true
	}
	return false
}
