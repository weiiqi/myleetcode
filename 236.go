package main

//最近公共祖先 TODO

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return nil
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
20210607
思路一： 递归
LCA具体在代码注释中
时间复杂度: O(n)
空间复杂度: O(n)
*/
func lowestCommonAncestor20210607(root, p, q *TreeNode) *TreeNode {
	//如果 PQ在一条线路上，则root节点就是最近的公共祖先
	if root == nil || root == p || root == q {
		return root
	}
	// 递归遍历左子树和右子树
	left := lowestCommonAncestor20210607(root.Left, p, q)
	right := lowestCommonAncestor20210607(root.Right, p, q)
	// 这里假设已经通过遍历找到对应的PQ节点，对具体情况进行判断

	// 左子树为空，说明PQ都不在左子树里面，二者可能都在右子树中，所以最近公共祖先是右节点
	if left == nil && right != nil {
		return right
	}
	// 右子树为空，说明PQ都不在右子树里面，二者可能都在左子树中，所以最近公共祖先是左节点
	if right == nil && left != nil {
		return left
	}
	// 如果左右子树都不为空，说明QP分散在左右子树中，那么root节点就是最近公共祖先
	if left != nil && right != nil {
		return root
	}
	// 左右子树都为空的情况，说明没有找到PQ，返回nil
	return nil
}
