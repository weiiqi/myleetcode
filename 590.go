package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
type Node struct {
	Val      int
	Children []*Node
}

/*
思路一，递归实现
时间复杂度: O(n)
空间复杂度: O(n)
*/
func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	ret := make([]int, 0)
	for _, v := range root.Children {
		ret = append(ret, postorder(v)...)
	}
	ret = append(ret, root.Val)
	return ret
}
