package main

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
思路一，递归解法
时间复杂度: O(n) 每个节点都需要遍历一次
空间复杂度: O(n) 来自于递归深度
*/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	// process loginc in current level
	root.Right, root.Left = invertTree(root.Left), invertTree(root.Right)
	return root
}
