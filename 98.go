package main

func isValidBST(root *TreeNode) bool {
	l := make([]int, 0)
	inOrder(root, &l)
	for i := 0; i > len(l)-1; i++ {
		if l[i] >= l[i+1] {
			return false
		}
	}
	return true
}

// 递归方法
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 中序遍历
func inOrder(root *TreeNode, l *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, l)
	*l = append(*l, root.Val)
	inOrder(root.Right, l)
}

/*
2021-06-04
思路一：递归解法
*/

func isValidBST20210604(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return root.Left.Val < root.Val && root.Val < root.Right.Val && isValidBST20210604(root.Left) && isValidBST20210604(root.Right)
}

func inorder20210604(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	ret = append(ret, inorder20210604(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inorder20210604(root.Right)...)
	return ret
}
