package main

func levelOrderBottom(root *TreeNode) [][]int {
	ret := levelOrder(root)
	for i,j := 0,len(ret)-1;i<j;i,j= i+1,j-1{
		ret[i],ret[j] = ret[j],ret[i]
	}
	return ret
}