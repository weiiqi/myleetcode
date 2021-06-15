package main

import "fmt"

/*
思路一： 分治法：分解问题
1. 2 ^ 5 = 2^2 * 2^2 * 2
2. 2^ 4 = 2^2 * 2^2
每次对n进行二分求，就能得到结果

时间复杂度: O(logn)
空间复杂度: O(1)
*/
func myPow(x float64, n int) float64 {
	if n < 0 {
		tmp := myPow_(x, -n)
		return 1 / tmp
	}
	return myPow_(x, n)
}

func myPow_(x float64, n int) float64 {
	// terminate
	if n == 0 {
		return 1
	}
	//process current level logic
	// drill down
	tmp := myPow_(x, n/2)
	// reverse state || merge result
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}

// 大牛写法
func pow_(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		return pow_(x*x, n/2)
	} else {
		return x * pow_(x*x, n/2)
	}
}

func main() {
	fmt.Println(myPow(2.0, 3))
}
