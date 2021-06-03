package main

import (
	"math"
)

// TODO，运行时间待优化
func maxPathSum(root *TreeNode) int {
	return scanTree(math.MinInt32, root)
}

func scanTree(max int, root *TreeNode) int {
	if root == nil {
		return max
	}
	rootValue := maxS(maxGrain(root.Left),0)+maxS(maxGrain(root.Right),0)+root.Val
	if max > rootValue{
		rootValue = max
	}
	return maxS(scanTree(rootValue, root.Left), scanTree(rootValue,root.Right))
}

func maxGrain(node *TreeNode) int {
	if node == nil{
		return 0
	}
	left := node.Val+maxS(maxGrain(node.Left),0)
	right := maxS(left,node.Val+ maxS(maxGrain(node.Right),0))
	return right
}

func maxS(a,b int) int {
	if a>b{
		return a
	}
	return b
}