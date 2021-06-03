package main

import "fmt"

/*
思路： 看最后一跳
状态：f(i) 表示是否能从0 到 i
推导： f(i) = OR(f(j), j<i && j能否跳到i) 判断之前所有的点是否能跳到当前节点
初始化：f(0) = 0
结果：f(n-1)
*/
func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == true && nums[j]+j >= i {
				dp[i] = true
			}
		}
	}
	return dp[len(nums)-1]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Print(canJump(nums))
}

func canJump1(nums []int) bool {
	n, jumpMax := len(nums), 0
	for i := 0; i < len(nums); i++ {
		if i <= jumpMax {
			jumpMax = max(jumpMax, i+nums[i])
			if jumpMax >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
