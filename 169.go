package main

import (
	"fmt"
	"sort"
)

/*
思路一：通过排序，下标n/2的即为众数
时间复杂度: O(n * log n)
空间复杂度: O(1)
*/

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

/*
思路二： hash表，记录每个元素出现的次数，返回次数超过n/2 的元素
时间复杂度: O(n)
空间复杂度: O(n)
*/

func majorityElement_2(nums []int) int {
	hashT := make(map[int]int)
	major := 0
	if len(nums)%2 == 0 {
		major = len(nums) / 2
	} else {
		major = (len(nums) / 2) + 1
	}
	for _, v := range nums {
		if _, ok := hashT[v]; ok {
			hashT[v]++
			if hashT[v] >= major {
				return v
			}
		} else {
			hashT[v] = 1
		}
	}
	return nums[0]
}

/*
思路三: 分治法
时间复杂度: O(n * logn)
空间复杂度: O(n * logn)
*/
func countInRange(nums []int, num, lo, hi int) int {
	count := 0
	for i := lo; i <= hi; i++ {
		if nums[i] == num {
			count++
		}
	}
	return count
}

func majortyElementRec(nums []int, lo, hi int) int {
	// terminate
	if lo == hi {
		return nums[lo]
	}
	// dril down
	// divide into 2 part
	mid := (hi-lo)/2 + lo
	left := majortyElementRec(nums, lo, mid)
	right := majortyElementRec(nums, mid+1, hi)

	if left == right {
		return left
	}
	// process cuurnt state
	// merge 2 part result
	leftCount := countInRange(nums, left, lo, hi)
	rightCount := countInRange(nums, right, lo, hi)

	// reverse state
	if leftCount > rightCount {
		return left
	}
	return right
}

func majortyElement_3(nums []int) int {
	return majortyElementRec(nums, 0, len(nums)-1)
}

/*
思路四：
时间复杂度: O(n)
空间复杂度: O(1)
*/

func majorityElement_4(nums []int) int {
	major := nums[0]
	count := 1
	for _, v := range nums {
		if count == 0 {
			count++
			major = v
		} else if major == v {
			count++
		} else {
			count--
		}
	}
	return major
}

func main() {
	nums := []int{8, 9, 8, 9, 8}
	fmt.Println(majorityElement_4(nums))
}
