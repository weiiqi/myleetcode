package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros3(nums)
	fmt.Println(nums)
}

// 个人想法 1
func moveZeroes(nums []int) {
	zidx := 0                        // 标记为0的元素
	for i := 0; i < len(nums); i++ { // 遍历数组
		if nums[i] != 0 { // 遇到非0的元素,数据交换
			nums[zidx] = nums[i]
			if i != zidx { // 下标和0标元素不相等的时候，则说明非0元素已经被交换，则需要将i的位置置为0
				nums[i] = 0
			}
			zidx++ // 0标志位后移
		}
	}
}

// 个人想法 2
// 开辟一个等长新数组，遍历第一个数组，每遇到一个非零数据，放到新数组中，后续位置补0, 即可

// 官方写法
func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

// 大佬写法

func moveZeros3(nums []int) {
	if len(nums) == 0 {
		return
	}
	insertPos := 0
	for _, v := range nums {
		if v != 0 {
			nums[insertPos] = v
			insertPos++
		}
	}
	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}
