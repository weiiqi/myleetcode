package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum20210528(nums))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	left, right := 1, len(nums)-1
	for first := 0; first < len(nums); first++ {
		if first >= 1 && nums[first] == nums[first-1] {
			continue
		}
		left = first + 1
		right = len(nums) - 1
		for left < right {
			if nums[first]+nums[left]+nums[right] < 0 {
				left++
			} else if nums[first]+nums[left]+nums[right] > 0 {
				right--
			} else {
				ret = append(ret, []int{nums[first], nums[left], nums[right]})
				left++
				right--
				for ; left < right; left++ {
					if nums[left] == nums[left-1] {
						continue
					} else {
						break
					}
				}
				for ; right > left; right-- {
					if nums[right] == nums[right+1] {
						continue
					} else {
						break
					}
				}
				continue
			}
		}
	}
	return ret
}

// 2020-05-27 结合 twoSum的灵感，尝试使用hash缓存结果的方式进行
// 思路1: a + b + c = 0 -> a + b = -c
// 先通过hash存储 每个值对应的hash
// 2021-05-28 这个方法有个很难的点，去重, 还是没想好方法，暂时搁置吧
func threeSum20210527(nums []int) [][]int {
	hasht := make(map[int]int)
	sort.Ints(nums)
	for _, v := range nums {
		hasht[v]++
	}
	ret := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			need := 0 - nums[i] - nums[j]
			cnt := 1
			if need == nums[i] {
				cnt++
			}
			if need == nums[j] {
				cnt++
			}
			if v, ok := hasht[need]; ok {
				if v >= cnt {
					ret = append(ret, []int{nums[i], nums[j], need})
					hasht[need] = v - cnt
				}
				continue
			}
		}
	}
	return ret
}

// 2021-05-28 尝试使用双指针方法
// 1. 对数组排序
// 2. 通过左右两个指针向中间偏移，left < right 循环下找符合条件的组合
func threeSum20210528(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[right] + nums[left]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				left++                       // 找到以后继续找
				right--                      // 找到以后继续找
				for ; left < right; left++ { // 剔除掉重复元素
					if nums[left] == nums[left-1] {
						continue
					}
					break
				}
				for ; left < right; right-- { // 剔除掉重复元素
					if nums[right] == nums[right+1] {
						continue
					}
					break
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
