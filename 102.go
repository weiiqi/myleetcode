package main

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}

/*
思路一: BFS
时间复杂度: O(n)
空间复杂度: O(n)
*/
func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	return nodeVal(&ret, root)
}
func nodeVal(ret *[][]int, nodes ...*TreeNode) [][]int {
	nextRootBode := make([]*TreeNode, 0)
	tmpNode := make([]int, 0)
	for _, v := range nodes {
		if v.Left != nil {
			nextRootBode = append(nextRootBode, v.Left)
		}
		if v.Right != nil {
			nextRootBode = append(nextRootBode, v.Right)
		}
		tmpNode = append(tmpNode, v.Val)
	}
	*ret = append(*ret, tmpNode)
	if len(nextRootBode) == 0 {
		return *ret
	}
	return nodeVal(ret, nextRootBode...)
}

/*
思路二：深度优先遍历
时间复杂度: O(n)
空间复杂度: O(n)
*/

func levelOrder_2(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}
	dfs(0, root, &ret)
	return ret
}

func dfs(level int, root *TreeNode, ret *[][]int) {
	// terminate
	if root == nil {
		return
	}
	// process current level logic
	if len(*ret) >= level {
		(*ret)[level] = append((*ret)[level], root.Val)
	} else {
		(*ret)[level] = []int{root.Val}
	}
	// drill down
	dfs(level+1, root.Left, ret)
	dfs(level+1, root.Right, ret)
}
