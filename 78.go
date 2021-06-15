package main

import (
	"fmt"
)

/*
思路一：分治法，类比归并排序
1. 拆分成小问题 , 对于子集来说，每个位置的情况就是选或者不选
2. 合并 就是递归出口的地方, 当遍历到最后列表最后一个元素时，就没有可以选的元素了
// 严格遵守递归的代码模版
时间复杂度: O(n * 2^n) -> 每个节点向下分散开，每个节点会被遍历2次，故为n^2.
空间复杂度: O(n * 2^n) -> 递归深度 + 临时存储结果的长度
*/
func subsets(nums []int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	subsets_(nums, false, 0, &path, &ret)
	return ret
}

func subsets_(nums []int, choose bool, index int, path *[]int, ret *[][]int) {
	//terminate
	if index == len(nums) {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}
	// process current level logic
	// drill down
	subsets_(nums, false, index+1, path, ret)
	if !choose && index > 0 && nums[index] == nums[index-1] {
		return
	}
	*path = append(*path, nums[index])

	subsets_(nums, true, index+1, path, ret)
	// reverse state  回溯需要，简单递归可能不需要
	*path = (*path)[:len(*path)-1]
	return
}

func subset_2(nums []int) [][]int {
	ret := make([][]int, 0)
	ret = append(ret, []int{})
	for i := 0; i < len(nums); i++ {
		tmpRet := make([][]int, len(ret))
		copy(tmpRet, ret)
		fmt.Println("before append tmpret")
		printSlicePtr(ret)

		fmt.Println("before range tmpret")
		printSlicePtr(tmpRet)

		for j := range tmpRet {
			tmpRet[j] = append(tmpRet[j], nums[i])
			if len(tmpRet) > 8 && j == 7 {
				fmt.Printf("%p\n", &tmpRet[7][3])
				fmt.Printf("%p\n", &tmpRet[15][3])
			}
		}
		fmt.Println("after range tmpret")
		printSlicePtr(tmpRet)

		ret = append(ret, tmpRet...)

		fmt.Println("after append tmpret")
		printSlicePtr(ret)
	}
	return ret
}

func printSlicePtr(nums [][]int) {
	for k, v := range nums {
		fmt.Printf("row-%v, len-%v, cap-%v, ptr-%p\n", k, len(v), cap(v), &nums[k])
		for i := 0; i < len(v); i++ {
			fmt.Printf("row-%v, clo-%v, ptr-%p\n", k, i, &v[i])
		}
	}
}

func main() {
	nums := []int{9, 0, 3, 5, 7}
	fmt.Println(subset_2(nums))
	//fmt.Println(subsets(nums))
}
