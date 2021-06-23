package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
思路一: BFS
时间复杂度：O(n)
空间复杂度：O(1)
*/
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		max := math.MinInt64
		for i := 0; i < l; i++ {
			if max < queue[i].Val {
				max = queue[i].Val
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
		}
		ret = append(ret, max)
		queue = queue[l:]
	}
	return ret
}
