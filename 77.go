package main

import "fmt"

/*
思路一： 看似简单，实则写不出来，看题解用回溯算法，但是回溯算法有点看不懂，啃一下
md 有点难
时间复杂度: O(n^k) * k
空间复杂度: O(n)
*/
func combine(n int, k int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)
	combine_(&path, 1, n, k, &ret)
	return ret
}

/*
递归的四步：
1. terminate
2. process current logic
3. drill down
4. reverse state
*/

func combine_(path *[]int, start, n, k int, ret *[][]int) {
	// 1
	if len(*path)+(n-start) < k-1 {
		return
	}
	if len(*path) == k {
		tmp := make([]int, len(*path))
		copy(tmp, *path)
		*ret = append(*ret, tmp)
		return
	}
	for i := start; i <= n; i++ {
		// 2
		*path = append(*path, i)
		// 3
		combine_(path, i+1, n, k, ret)
		// 4
		*path = (*path)[:len(*path)-1]
	}
	return
}

func main() {
	ret := combine(3, 3)
	fmt.Println(ret)
}
