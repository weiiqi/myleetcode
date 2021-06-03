package main

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left+1 < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}

	if nums[left] >= target {
		return left
	} else if nums[right] >= target {
		return right
	} else if nums[right] < target {
		return right + 1
	}
	return 0
}
