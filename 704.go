package main

import "fmt"

func search(nums []int, target int) int {
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		mid := (start + end) / 2
		if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		} else {
			return mid
		}
	}
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Print(search(nums, 0))
}
