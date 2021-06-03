package main

import (
	"container/heap"
	"fmt"
)

func main()  {
	matrix := [][]int{
		[]int{1,5,9},
		[]int{10,11,13},
		[]int{12,13,15},
	}
	fmt.Print(kthSmallest(matrix, 8))
}

func kthSmallest(matrix [][]int, k int) int {
	h := new(heapNode)
	for _,v := range matrix{
		for _,sv:= range(v){
			if len(*h) < k{
				heap.Push(h,sv)
			}else{
				if (*h)[0] > sv{
					heap.Pop(h)
					heap.Push(h,sv)
				}
			}
		}
	}
	return heap.Pop(h).(int)
}


type heapNode []int

func (h *heapNode)Less(i,j int) bool  {
	return (*h)[i] > (*h)[j]
}

func (h *heapNode)Len() int  {
	return len(*h)
}

func (h *heapNode)Swap(i,j int)  {
	(*h)[i],(*h)[j] = (*h)[j],(*h)[i]
}

func (h *heapNode)Push(x interface{}) {
	(*h) = append((*h), x.(int))
}

func (h *heapNode)Pop() interface{}  {
	old := *h
	n := len(*h)
	x := (*h)[n-1]
	*h = old[:n-1]
	return x
}
