package main

// TODO 思路: 实在不行就用笨方法，hashtable存放次数，然后返回次数为1的
// TODO 异或位运算
func singleNumber2(nums []int) int {
	var ret int
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		ret ^= (sum % 3) << i
	}
	return ret
}
