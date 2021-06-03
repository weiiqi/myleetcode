package main

import "fmt"

// 汉明距离 思路：TODO 可背
func hammingWeight(num uint32) int {
	res := 0
	for num != 0 {
		num = num & (num - 1)
		res++
	}
	return res
}

func main() {
	fmt.Print(hammingWeight(5))
}
