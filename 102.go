package main

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}

func levelOrder(root *TreeNode) [][]int {
    ret :=make([][]int,0)
    if root == nil{
        return ret
    }
    return nodeVal(&ret, root)
}
func nodeVal(ret *[][]int, nodes ...*TreeNode) [][]int  {
    nextRootBode := make([]*TreeNode,0)
    tmpNode := make([]int, 0)
    for _,v := range nodes{
        if v.Left != nil{
            nextRootBode = append(nextRootBode, v.Left)
        }
        if v.Right != nil{
            nextRootBode = append(nextRootBode,v.Right)
        }
        tmpNode = append(tmpNode, v.Val)
    }
    *ret = append(*ret, tmpNode)
    if len(nextRootBode) == 0{
        return *ret
    }
    return nodeVal(ret, nextRootBode...)
}

