package main

import "math"

type MinStack struct {
	min   []int
	stack []int
}

///** initialize your data structure here. */
//func Constructor() MinStack {
//	return MinStack{
//		min:   make([]int, 0),
//		stack: make([]int, 0),
//	}
//}
//
//func (this *MinStack) Push(x int) {
//	min := this.GetMin()
//	if x < min {
//		this.min = append(this.min, x)
//	} else {
//		this.min = append(this.min, min)
//	}
//	this.stack = append(this.stack, x)
//}
//
//func (this *MinStack) Pop() {
//	if len(this.stack) == 0 {
//		return
//	}
//	this.stack = this.stack[:len(this.stack)-1]
//	this.min = this.min[:len(this.min)-1]
//}
//
//func (this *MinStack) Top() int {
//	if len(this.stack) == 0 {
//		return 0
//	}
//	return this.stack[len(this.stack)-1]
//}
//
//func (this *MinStack) GetMin() int {
//	if len(this.min) == 0 {
//		return math.MaxInt32
//	}
//	min := this.min[len(this.min)-1]
//	return min
//}
//*/
//
///**
// * Your MinStack object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * obj.Pop();
// * param_3 := obj.Top();
// * param_4 := obj.GetMin();
// */

/*
2021-06-02
思路： 使用两个栈来实现
1. 一个栈用来存放本来的数据顺序
2. 一个栈用来存放当前的最小值
*/

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: make([]int, 0),
		min:   make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.GetMin() < val {
		this.min = append(this.min, this.GetMin())
	} else {
		this.min = append(this.min, val)
	}
	return
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[:len(this.stack)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		// 这里是个细节，因为前面push时候会跟这个值做对比,所以这里默认要是个最大值
		// 或者在push时候判空处理
		return math.MaxInt32
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 *
 */
