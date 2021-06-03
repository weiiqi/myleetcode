package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil{
		return make([][]int, 0)
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ret := make([][]int, 0)
	direct := true
	for len(queue) > 0{
		list := make([]int, 0)
		l := len(queue)
		for i:=0; i<l;i++{
			level := queue[0]
			queue = queue[1:]
			list = append(list, level.Val)
			if level.Right != nil{
				queue = append(queue, level.Right)
			}
			if level.Left != nil{
				queue = append(queue, level.Left)
			}
		}
		if direct{
			list = reverseL(list)
		}
		direct = !direct
		ret = append(ret, list)
	}
	return ret
}

func reverseL(nums []int) []int {
	for i,j:= 0,len(nums)-1;i<j;i,j=i+1,j-1{
		nums[i],nums[j]=nums[j],nums[i]
	}
	return nums
}