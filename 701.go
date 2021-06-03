package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	tmpNode :=root
	newNode := &TreeNode{Val:val}
	if tmpNode == nil{
		return newNode
	}
	for tmpNode != nil{
		if val > tmpNode.Val{
			if tmpNode.Right == nil{
				tmpNode.Right = newNode
				break
			}else{
				tmpNode = tmpNode.Right
			}
		}else{
			if tmpNode.Left == nil{
				tmpNode.Left = newNode
				break
			}else{
				tmpNode = tmpNode.Left
			}
		}
	}
	return root
}