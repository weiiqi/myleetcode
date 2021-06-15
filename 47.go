package main

import (
	"fmt"
	"sort"
)

/*
思路一：结合46题，这类的题都适用回溯算法
区别于46题，
1. 需要去除重复的元素，先用排序，保证相同的数据相邻
2. 在每次for循环遍历的时候，判断下是不是和前一个元素相同，且已经被访问过
3. 如果访问过了，就跳过继续下一个元素的遍历
时间复杂度: O(n * n!)
空间复杂度: O(n)  需要标记是否已经访问过，n + 递归深度n = O(2n) -> O(n)
*/
func permuteUnique(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	visited := make([]bool, len(nums))
	sort.Ints(nums)
	permuteUnique_(nums, &path, visited, &ret)
	return ret
}

func permuteUnique_(nums []int, path *[]int, visited []bool, ret *[][]int) {
	if len(*path) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] && visited[i-1] {
			continue
		}
		*path = append(*path, nums[i])
		visited[i] = true
		permuteUnique_(nums, path, visited, ret)
		visited[i] = false
		*path = (*path)[:len(*path)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permuteUnique(nums))
}
