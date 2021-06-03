package main

func singleNumber(nums []int) int {
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = tmp | nums[i]
	}
	return tmp
}
