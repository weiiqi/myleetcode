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
2021-06-07
思路一：前序遍历-> 根左右， 中序遍历 -> 左根右。前序确定了根的位置，中序确定了左右子树的位置和元素个数。
递归出口：当前序遍历完毕时，说明根节点遍历完毕，递归结束
时间复杂度: O(n)
空间复杂度: O(n)
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0], Left: nil, Right: nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// len(inorder[:i])
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	root := buildTree(pre, in)
	fmt.Println(root)
}
