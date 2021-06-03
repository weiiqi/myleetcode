package main

import (
	"container/heap"
	"math"
	"sort"
)

/*
思路一：
1. 计算出窗口每次的起始位置
2. 记录下窗口内的最大值
时间复杂度： O(k*n)
空间复杂度： O(1)
*/
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) <= k {
		return []int{max(nums)}
	}
	ret := make([]int, 0)
	for i := 0; i+k < len(nums); i++ {
		ret = append(ret, max(nums[i:i+k]))
	}
	return ret
}

func max(l []int) int {
	max := math.MinInt32
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

/*
思路二：
借鉴最小栈的思想，遍历一次然后拿到数组对应的最大栈
返回 nums[k-1:],即时窗口的最大值集合
时间复杂度： O(n)
空间复杂度： O(k)
*/

// 这种写法忽略了最大值移除窗口的情况, 但是这种单调思维是对的
func maxSlidingWindow_1(nums []int, k int) []int {
	maxStack := []int{math.MinInt32}
	for _, v := range nums {
		if v > maxStack[len(maxStack)-1] {
			maxStack = append(maxStack, v)
		} else {
			maxStack = append(maxStack, maxStack[len(maxStack)-1])
		}
	}
	return maxStack[k:]
}

// 考虑窗口左右边界以及最大值在窗口外的情况
/*
思路2.1:
1. 建立单调队列
2. 遍历元素，保证最大值在队头，如果在窗口内，将队头元素放入结果队列, 更新窗口边界
3. 如果队头元素不在窗口内, 移除队头元素，继续遍历数组
*/
func maxSlidingWindow_1_1(nums []int, k int) []int {
	maxStack, ret := make([]int, 0), make([]int, 0)
	// left 窗口左边届，right 窗口右边界
	for left, right := 0, 0; right < len(nums); right++ {
		if len(maxStack) == 0 {
			maxStack = append(maxStack, right)
			continue
		}
		// 保证栈底为最大值
		for len(maxStack) > 0 && nums[right] > nums[maxStack[len(maxStack)-1]] {
			maxStack = maxStack[:len(maxStack)-1]
		}
		maxStack = append(maxStack, right)
		if right-left+1 == k {
			for maxStack[0] < left {
				maxStack = maxStack[1:]
			}
			ret = append(ret, nums[maxStack[0]])
			left++
		}
	}
	return ret
}

/*
参考官方题解，思路二属于单调队列解法
思路三，大顶堆解法
1. 借助heap interface 实现大顶堆 swap
2. 遍历数组，遇到元素就丢进heap中，堆顶为最大值
3. 删除窗口外的值
4. 重点考虑下窗口边界
时间复杂度: O(n*logn)
空间复杂度: O(n)
*/

type bigHeap struct {
	nums []int
	heap []int
}

func (bh *bigHeap) Len() int {
	return len(bh.heap)
}

func (bh *bigHeap) Swap(i, j int) {
	bh.heap[i], bh.heap[j] = bh.heap[j], bh.heap[i]
}

func (bh *bigHeap) Less(i, j int) bool {
	return bh.nums[bh.heap[i]] > bh.nums[bh.heap[j]]
}

// 堆顶就是nums中最大的元素
func (bh *bigHeap) Pop() interface{} {
	tmp := bh.heap[len(bh.heap)-1]
	bh.heap = bh.heap[:len(bh.heap)-1]
	return tmp
}

func (bh *bigHeap) Push(v interface{}) {
	bh.heap = append(bh.heap, v.(int))
}

func maxSlidingWindow_2(nums []int, k int) []int {
	bh := &bigHeap{nums: nums, heap: make([]int, k)}
	ret := make([]int, 0)
	heap.Init(bh)
	for left, right := 0, 0; right < len(nums); right++ {
		heap.Push(bh, right)
		if right-left+1 == k {
			for left > bh.heap[0] {
				heap.Pop(bh)
			}
			ret = append(ret, nums[bh.heap[0]])
			left++
		}
	}
	return ret
}
