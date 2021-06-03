package main

import "fmt"

/*
思路：
找到由左向右遍历，遇到重复的就通过切片进行删除
然后返回数组len值
*/
func removeDuplicates(nums []int) int {
	i := 0
	for i < len(nums) {
		if i > 0 && nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return i
}

/*
思路二：
双指针, 快慢指针方式
*/

func removeDuplicates_1(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
func main() {
	nums := []int{1, 1, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates_1(nums))
	fmt.Println(nums)
}
