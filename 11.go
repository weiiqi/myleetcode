package main

import "fmt"

func main() {
	height := []int{4, 3, 2, 1, 4}
	fmt.Println(maxArea2(height))
}

/*
/*  个人思路 O(N^2) 两层遍历，分别计算出面积，记录最大的返回
*/
func maxArea(height []int) int {
	area := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			len := j - i
			wide := height[j]
			if height[i] < height[j] {
				wide = height[i]
			}
			tmpArea := len * wide
			if tmpArea > area {
				area = tmpArea
			}
		}
	}
	return area
}

/*
   尝试O(N)解法，双指针，左右指针，往中间聚拢，计算面积，比较返回最大面积
*/

func maxArea2(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		tmp := 0
		if height[right] < height[left] {
			tmp = height[right] * (right - left)
			right--
		} else {
			tmp = height[left] * (right - left)
			left++
		}
		if tmp > maxArea {
			maxArea = tmp
		}
	}
	return maxArea
}
