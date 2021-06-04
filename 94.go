package main

// 思路： 通过stack， 保存已经访问的元素，用于原路返回
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 优先左节点入栈
		}
		// 数据弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, val.Val)
		root = val.Right
	}
	return ret
}

// 递归思路，简便。但是内存消耗比较大
func inorder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, inorder(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorder(root.Right)...)
	return ret
}

/*
2021-06-03
思路一：递归访问
递归会使用更多的内存资源，因为除去本身需要保存的值，还需要保存程序运行时的状态值
时间复杂度: O(n)
空间复杂度: O(n)
*/

// 中序遍历
func inorderTraversal20210603(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, inorderTraversal20210603(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorderTraversal20210603(root.Right)...)
	return ret
}

// 前序遍历
func preorderTraversal20210603(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, root.Val)
	ret = append(ret, preorderTraversal20210603(root.Left)...)
	ret = append(ret, preorderTraversal20210603(root.Right)...)
	return ret
}

// 后序遍历
func postorderTraversal20210603(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, postorderTraversal20210603(root.Left)...)
	ret = append(ret, postorderTraversal20210603(root.Right)...)
	ret = append(ret, root.Val)
	return ret
}
