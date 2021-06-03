package main

import (
	"container/heap"
	"fmt"
)

// TOPN

func main()  {
	nums := []int{5,2,5,3,5,3,1,1,3}
	fmt.Print(topKFrequent(nums, 2))
}

// 小顶堆 做法
func topKFrequent(nums []int, k int) []int {
	hashM := make(map[int]int)
	for _,v := range nums{
		_, ok := hashM[v]
		if !ok{
			hashM[v] = 1
		}else{
			hashM[v]++
		}
	}
	h := new(nodeHeap)
	topK := min(k,len(nums))
	size := 0
	for k,v := range hashM{
		if size < topK{
			heap.Push(h, &Node{val:k,times:v})
			size++
		}else{
			if v > (*h)[0].times{
				heap.Pop(h)
				heap.Push(h,&Node{val:k,times:v})
			}
		}
	}
	ret := make([]int, topK)
	for i:=topK-1;i>=0;i--{
		ret[i] = heap.Pop(h).(*Node).val
	}
	return ret
}

type Node struct {
	val int
	times int
}

type nodeHeap []*Node

func (h nodeHeap)Len() int  {
	return len(h)
}

func (h nodeHeap)Less(i,j int) bool {
	return h[i].times < h[j].times
}

func (h nodeHeap)Swap(i,j int)  {
	h[i],h[j] = h[j],h[i]
}

func (h *nodeHeap)Push(x interface{})  {
	*h = append(*h, x.(*Node))
}

func (h *nodeHeap)Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func min(a,b int) int {
	if a< b{
		return a
	}
	return b
}

//TODO 大顶堆