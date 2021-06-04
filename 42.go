package main

import "fmt"

/*
思路一：
灵感来自于84题，用单调栈来解。该题本身是在寻找每个柱子左右两边的第一个最大值
由左向右遍历，把每根柱子的长短记录下来，保证栈底的元素最大
遇到入栈元素比栈顶元素大的情况，就可以计算该元素的能接受的最大雨水

时间复杂度: O(n)
空间复杂度: O(n)
*/

func trap(height []int) int {
	ret := 0
	stack := make([]int, 0)
	for k, v := range height {
		for len(stack) > 0 && v >= height[stack[len(stack)-1]] {
			if len(stack) < 2 {
				stack = stack[:len(stack)-1]
				continue
			}
			area := (k - stack[len(stack)-2] - 1) * (min(v, height[stack[len(stack)-2]]) - height[stack[len(stack)-1]])
			ret = ret + area
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, k)
	}
	return ret
}

/*
思路二：学习国际站的大佬写法
计算面积的方式不再是 area = wide * length
修改为累计的方式(宽度固定为1)， 只要确定左右两个最高的柱子，那么中间的值就是柱子减去柱子高度，累加起来就是阴影部分的面积
难思考，但是是个好的方式。
时间复杂度: O(n)
空间复杂度: O(1)
*/

func trap_2(height []int) int {
	left, right := 0, len(height)-1
	maxleft, maxright := 0, 0
	ret := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] > maxleft {
				maxleft = height[left]
			} else {
				ret += maxleft - height[left]
			}
			left++
		} else {
			if height[right] > maxright {
				maxright = height[right]
			} else {
				ret += maxright - height[right]
			}
			right--
		}

	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Print(trap_2(nums))
}
