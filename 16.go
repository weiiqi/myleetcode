package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := make([]int,0)
	for first:=0;first<len(nums);first++{
		left := 0
		right := len(nums)-1
		for left < first && right >first{
			total := nums[first] + nums[left] + nums[right]
			ret = append(ret, total-1)
		}
	}
}
