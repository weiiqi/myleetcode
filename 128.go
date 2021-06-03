package main

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, v := range nums {
		numSet[v] = true
	}
	lonest := 0
	for num := range numSet {
		if !numSet[num-1] {
			currNum := num
			currStreak := 0
			for numSet[currNum+1] {
				currNum++
				currStreak++
			}
			if lonest < currStreak {
				lonest = currStreak
			}
		}
	}
	return lonest
}
