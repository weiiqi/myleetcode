package main

func findMin2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		mid := start + (end-start)/2
		if nums[mid] <= nums[end] {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] < nums[end] {
		return nums[start]
	}
	return nums[end]
}
