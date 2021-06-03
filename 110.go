package main

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil{
		return true
	}
	if root.Left == nil && root.Right == nil{
		return true
	}
	leftH := maxDepth(root.Left)-1
	rightH := maxDepth(root.Right)-1
	if math.Abs(float64(leftH-rightH)) > 1{
		return false
	}else{
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}
