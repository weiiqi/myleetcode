package main

func maxDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if root.Right == nil && root.Left == nil{
		return 1
	}
	l := maxDepth(root.Left)
	r := maxDepth(root.Right)
	if l >r {
		return l+1
	}else{
		return r+1
	}
}