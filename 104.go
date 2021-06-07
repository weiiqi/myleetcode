package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Right == nil && root.Left == nil {
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

/*
2021-06-07
思路一： 递归
最大深度= 左子树/右子树深度 + 1
时间复杂度: O(n)
空间复杂度: O(height) height 树的深度
*/

func maxDepth20210607(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth20210607(root.Left)
	right := maxDepth20210607(root.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}

}
