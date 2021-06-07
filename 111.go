package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路一：本质是在求最小深度，与求最大深度类似结构
时间复杂度: O(n)
空间复杂度: O(height) 二叉树深度
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 || right == 0 {
		return left + right + 1
	}
	if left < right {
		return left + 1
	}
	return right + 1
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node1.Left = node2
	fmt.Println(minDepth(node1))
}
