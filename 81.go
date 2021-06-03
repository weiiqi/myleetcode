package main

// 搜索旋转排序数组
func search2(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	for start+1 < end {
		for start < end && nums[start] == nums[start+1] {
			start++
		}
		for start < end && nums[end] == nums[end-1] {
			end--
		}
		mid := start + (end-start)/2
		if nums[mid] == target {
			return true
		}
		if nums[start] < nums[mid] {
			if nums[start] <= target && target <= nums[mid] {
				end = mid
			} else {
				start = mid
			}
		} else if nums[end] > nums[mid] {
			if nums[mid] <= target && target <= nums[end] {
				start = mid
			} else {
				end = mid
			}
		}
	}
	if nums[start] == target || nums[end] == target {
		return true
	}
	return false
}

//二分法
//1、初始化：start=0、end=len-1
//2、循环退出条件：start + 1 < end
//3、比较中点和目标值：A[mid] ==、 <、> target
//4、判断最后两个元素是否符合：A[start]、A[end] ? target
