package main

// 切片递归
func sortedArrayToBST1(nums []int) *TreeNode {
	if len(nums) <=0{
		return nil
	}
	rootIdx := len(nums)/2
	root := new(TreeNode)
	root.Val = nums[rootIdx]
	root.Left = sortedArrayToBST1(nums[:rootIdx])
	root.Right = sortedArrayToBST1(nums[rootIdx+1:])
	return root
}


// idx 递归
func sortedArrayToBST2(nums []int) *TreeNode {
	return helper1(nums, 0, len(nums)-1)
}

func helper1(nums []int, left, right int) *TreeNode {
	if left > right{
		return nil
	}
	rootIdx := (left+right)/2
	root := new(TreeNode)
	root.Val = nums[rootIdx]
	root.Left = helper1(nums, left, rootIdx+1)
	root.Right = helper1(nums, rootIdx-1, right)
	return root
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
