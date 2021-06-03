package main

// 思路： TODO
func updateMatrix(matrix [][]int) [][]int {
	q := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				point := []int{i, j}
				q = append(q, point)
			} else {
				matrix[i][j] = -1
			}
		}
	}
	//directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	//for len(q) != 0 {
	//	point := q[0]
	//	q = q[1:]
	//	for _, v := range directions {
	//		x := point[0] + v[0]
	//		y := point[1] + v[1]
	//	}
	//}
	return matrix
}
