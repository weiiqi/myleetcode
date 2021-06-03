package main

import "fmt"

func singleNumber3(nums []int) []int {
	diff := 0
	for i := 0; i < len(nums); i++ {
		diff ^= nums[i]
	}
	ret := []int{diff, diff}
	diff = (diff & (diff - 1)) ^ diff
	for i := 0; i < len(nums); i++ {
		if diff&nums[i] == 0 {
			ret[0] ^= nums[i]
		} else {
			ret[1] ^= nums[i]
		}
	}
	return ret
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Print(singleNumber3(nums))
}
