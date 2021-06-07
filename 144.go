package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
思路一：使用递归方式进行前序遍历，相对栈循环的方式，这样思路较清晰
时间复杂度: O(n)
空间复杂度: O(n)
*/
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal(root.Left)...)
	ret = append(ret, preorderTraversal(root.Right)...)
	return ret
}

/*
思路二：通过遍历来实现(来自官方题解)
时间复杂度: O(n)
空间复杂度: O(n)
*/

func preorderTraversal20210604(root *TreeNode) []int {
	stack := []*TreeNode{}
	node := root
	ret := make([]int, 0)
	for node != nil || len(stack) > 0 {
		for node != nil {
			ret = append(ret, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return ret
}
