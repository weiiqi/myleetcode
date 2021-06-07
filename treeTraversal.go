package main

/*
由于二叉树的迭代遍历一直没有思路，三种方式记录在这里
*/
func preorderTraversal(root *TreeNode) []int {
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

func inorderTraversal20200604(root *TreeNode) []int {
	stack := []*TreeNode{}
	ret := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, root.Val)
		root = root.Right
	}
	return ret
}

func postorderTraversal(root *TreeNode) {
	stack := []*TreeNode{}
	var prev *TreeNode
	ret := make([]int, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Right == nil || root.Right == prev {
			ret = append(ret, root.Val)
			prev = root
			root = nil
		} else {
			stack = append(stack, root)
			root = root.Right
		}
	}
	return ret
}
