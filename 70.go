package main

import "fmt"

// 递归思路, 超出时间限制
func climbStairs(n int) int {
	funcSet := make(map[int]int)
	stepn_1, stepn_2 := 0, 0
	if v, ok := funcSet[n-1]; ok {
		stepn_1 = v
	} else {
		stepn_1 = step(n - 1)
		funcSet[n-1] = stepn_1
	}
	if v, ok := funcSet[n-2]; ok {
		stepn_2 = v
	} else {
		stepn_2 = step(n - 2)
		funcSet[n-2] = stepn_2
	}
	return stepn_1 + stepn_2
}

func step(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + step(n-2)
}

// 思路： 动态规划
func climbStairs1(n int) int {
	if n == 1 || n == 0 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs_20210527(4))
	fmt.Println(climbStairs_20210527_1(4))
}

// 2021-5-27
// 思路 1 ： 递归  注意缓存各步骤的计算结果，防止重复计算，空间换时间的计算方法
func climbStairs_20210527(n int) int {
	cache := make(map[int]int)
	return step_20210527(n, cache)
}

func step_20210527(n int, cache map[int]int) int {
	if n <= 2 {
		return n
	}
	if v, ok := cache[n]; ok {
		return v
	} else {
		stepn := step_20210527(n-1, cache) + step_20210527(n-2, cache)
		cache[n] = stepn
		return stepn
	}
}

// 2021-05-27
// 思路二： 动态规划
// 递推公式： dp[x] = dp[x-1] + dp[x-2]
// 这里不需要缓存结果，只需要最终结果，所以只需要数字即可

func climbStairs_20210527_1(n int) int {
	if n <= 2 {
		return n
	}
	pre, prep := 1, 0
	next := 0
	for i := 0; i < n; i++ {
		next = pre + prep
		prep = pre
		pre = next
	}
	return next
}
