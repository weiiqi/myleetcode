package main

import "fmt"

/*
思路一：递归，来自于77题的自信，使用回溯法来解决问题
// 来自于题解的做法
时间复杂度: O(n * n!)
空间复杂度: O(n * n!)
*/
func permute(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	visited := make([]bool, len(nums), cap(nums))
	permute_(nums, &path, &ret, visited)
	return ret
}

/*
1. 从剩下的数字中选出一个数字来，然后递归选取数字
*/

/*
2. 这里思路有个误区，存放剩余可用节点，但是这给递归带来了麻烦，而且有点绕, 通过hash解决了这个问题挺好
*/
func permute_(nums []int, path *[]int, ret *[][]int, visited []bool) {
	// terminate
	if len(nums) == len(*path) {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
	}
	// process current level logic
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		*path = append(*path, nums[i])
		visited[i] = true
		// drill down
		permute_(nums, path, ret, visited)
		// reverse state
		*path = (*path)[:len(*path)-1]
		visited[i] = false
	}

}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permute(nums))
}
