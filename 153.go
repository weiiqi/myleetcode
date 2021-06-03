package main

import "fmt"

// 思路：最后一个值作为target， 然后往左移动，最后比较start,end的值
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := start + (end-start)/2
		// 最后一个元素为target
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] > nums[end] {
		return nums[end]
	}
	return nums[start]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	fmt.Print(findMin(nums))
}
