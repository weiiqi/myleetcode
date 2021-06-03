package main

import "fmt"

// 搜索旋转排序列表
// 思路：二分法，四种情况判断 TODO
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		// 相等直接返回
		if nums[mid] == target {
			return mid
		}
		// 判断在哪个区间，可能分为四种情况
		if nums[start] < nums[mid] { // 左半部分
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] { // 右半部分
			if nums[end] >= target && nums[mid] <= target {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target {
		return start
	} else if nums[end] == target {
		return end
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Print(search1(nums, 6))
}
