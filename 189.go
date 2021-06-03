package main

import "fmt"

/*
思路一：
如果k和len(nums)求余等于0，则返回原数组
实现一个右移一位的小函数，然后循环K次，每次移动一位
超时了: 应该是因为元素移动的次数太多，在数组很大的时候，这个问题就会被放大
时间复杂度：O(n^k)
空间复杂度：O(1)
*/
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	t := k % len(nums)
	if t == 0 {
		return
	}
	for i := 0; i < t; i++ {
		rotateOne(nums)
	}

}

func rotateOne(nums []int) {
	end := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		nums[i+1] = nums[i]
	}
	nums[0] = end
}

/*
思路二：
开辟新的元素数组 然后通过copy把新数据迁移至nums中
时间复杂度：O(N)
空间复杂度：O(N)
*/

func rotate_2(nums []int, k int) {
	if len(nums) < 2 {
		return
	}
	t := k % len(nums)
	if t == 0 {
		return
	}
	newNums := make([]int, 0)
	newNums = append(newNums, nums[len(nums)-t:]...)
	newNums = append(newNums, nums[:len(nums)-t]...)
	copy(nums, newNums)
	return
}

/*
思路三: 看了刷友的题解，真是人才
1. 反转全部数组
2. 反转nums[:k]
3. 反转nums[k:]
时间复杂度: O(n)
空间复杂度：O(1)
*/

func reverseArray(nums []int) {
	for left, right := 0, len(nums)-1; left < right; {
		nums[left], nums[right] = nums[right], nums[left]
		right--
		left++
	}
	return
}

func rotate_3(nums []int, k int) {
	k %= len(nums)
	reverseArray(nums)
	reverseArray(nums[:k])
	reverseArray(nums[k:])
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rotate_3(nums, 3)
	fmt.Println(nums)
}
