package main

import "fmt"

/*
思路一：
*/
type MyCircularDeque struct {
	queue []int
	cap   int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: make([]int, 0),
		cap:   k,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(vgalue int) bool {
	if this.IsFull() {
		return false
	}
	this.queue = append(this.queue, vgalue)
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue = append([]int{value}, this.queue[:]...)
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[:len(this.queue)-1]
		return true
	}
	return false
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if !this.IsEmpty() {
		this.queue = this.queue[1:]
		return true
	}
	return false
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if !this.IsEmpty() {
		return this.queue[len(this.queue)-1]
	}
	return -1
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if !this.IsEmpty() {
		return this.queue[0]
	}
	return -1
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.queue) == 0
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.cap <= len(this.queue)
}

/*
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

func main() {
	obj := Constructor(3)
	obj.InsertLast(1)
	obj.InsertLast(2)
	obj.InsertFront(4)
	fmt.Println(obj.GetFront())
	obj.DeleteFront()
	fmt.Println(obj)
}
