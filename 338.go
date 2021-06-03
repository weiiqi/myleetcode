package main

import "fmt"

func countBits(num int) []int {
	ret := make([]int, 0)
	for i := 0; i <= num; i++ {
		ret = append(ret, count(i))
	}
	return ret
}

func count(num int) int {
	ret := 0
	for num != 0 {
		num = num & (num - 1)
		ret++
	}
	return ret
}

// 思想 动态规划
func countBits1(num int) []int {
	res := make([]int, num+1)
	for i := 1; i <= num; i++ {
		// 上一个缺1的元素+1即可
		res[i] = res[i&(i-1)] + 1
	}
	return res
}

func main() {
	fmt.Print(countBits1(5))
}
