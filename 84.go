package main

import "fmt"

//func largestRectangleArea(heights []int) int {
//	min := 0
//	max := len(heights) - 1
//	ret := 0
//	for i := 0; i < max; i++ {
//		left := i
//		right := i
//		for left > min && right < max {
//			if heights[left] >= heights[i] {
//				left--
//				continue
//			}
//			if heights[right] >= heights[i] {
//				right++
//				continue
//			}
//		}
//		if tmp := heights[i] * (right - left); tmp > ret {
//			ret = tmp
//		}
//	}
//	return ret
//}

// 思路，重点是查找i 左右两边的值，算出宽度，计算出面积
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	max := 0
	for i := 0; i <= len(heights); i++ {
		var cur int
		if i == len(heights) {
			cur = 0
		} else {
			cur = heights[i]
		}
		for len(stack) != 0 && cur <= heights[stack[len(stack)-1]] {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			h := heights[pop]
			w := i
			if len(stack) != 0 {
				peek := stack[len(stack)-1]
				w = i - peek - 1
			}
			max = Max(max, h*w)
		}
		stack = append(stack, i)
	}
	return max
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

/*
2021-06-02 暴力解法
按照柱子高度，寻找比柱子小的最近的左右边界，则就是该柱子能组成的最大阴影面积
时间复杂度: O(n^2)
空间复杂度: O(1)
*/
func largestRectangleArea20210602(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		left, right := i, i
		for left-1 >= 0 {
			if heights[left] < heights[i] {
				break
			}
			left--
		}
		for right+1 <= len(heights) {
			if heights[right] < heights[i] {
				break
			}
			right++
		}
		area := (right - left + 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

/*
2021-06-02
思路二： 学习官方题解，单调栈方式
对比上述暴力方式，主要优化点，在寻找柱子左右边界(第一个小于该柱子高度的柱子)。在由左向右遍历的过程中，
每个柱子的左下标是可以被确定下来的, 剩下的只需要在向右遍历的过程中，确定右下标，即可确定该柱子能组成的最大面积
时间复杂度: O(n)
空间复杂度: O(n)
*/
func largestRectangleArea20210602_1(heights []int) int {
	maxArea := 0
	stack := []int{-1}
	for k, v := range heights {
		for stack[len(stack)-1] != -1 && v < heights[stack[len(stack)-1]] {
			area := heights[stack[len(stack)-1]] * (k - (stack[len(stack)-2]) - 1)
			if area > maxArea {
				maxArea = area
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}
	for len(stack) > 1 {
		area := heights[stack[len(stack)-1]] * (len(heights) - stack[len(stack)-2] - 1)
		if area > maxArea {
			maxArea = area
		}
		stack = stack[:len(stack)-1]
	}
	return maxArea
}

func main() {
	nums := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea20210602_1(nums))
}
