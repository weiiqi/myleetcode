package main

import "sort"

/*
思路一：
开辟新的切片，归并排序的方式组成新的数组
时间复杂度：O(n)
空间复杂度: O(n)
*/

func mergeSortedList(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	tmp := make([]int, 0)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			tmp = append(tmp, nums1[i])
			i++
		} else {
			tmp = append(tmp, nums2[j])
			j++
		}
	}
	if i < m {
		tmp = append(tmp, nums1[i:]...)
	}
	if j < n {
		tmp = append(tmp, nums2[j:]...)
	}
	copy(nums1, tmp)
	return
}

/*
思路二：先合并再排序
官方题解
时间复杂度：O(nlogn)
空间复杂度：O(logn)
对应快排的时间和空间复杂度
*/

func merge_2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	return
}

/*
思路三：逆向双指针(官方题解)
刚开始增尝试使用正向双指针方式实现，但是会产生元素覆盖的情况，放弃了, 看到逆向双指针，一拍大腿真tm秒
时间复杂度: O(n)
空间复杂度: O(1)
*/

func merge_3(nums1 []int, m int, nums2 []int, n int) {
	p1, p2, tail := m-1, n-1, m+n-1
	for tail >= 0 {
		if p1 < 0 {
			nums1[tail] = nums2[p2]
			p2--
			tail--
		}
		if p2 < 0 {
			nums1[tail] = nums1[p1]
			p1--
			tail--
		}
		if nums2[p2] > nums1[p1] {
			nums1[tail] = nums2[p2]
			tail--
			p2--
		} else {
			nums1[tail] = nums1[p1]
			p1--
			tail--
		}
	}
	return
}
