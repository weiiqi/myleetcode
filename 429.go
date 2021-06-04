package main

import "fmt"

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
思路一：
维护一个节点列表，根节点在第一层，直到节点列表为空为止
时间复杂度: O(n) 每个节点只被遍历一次
空间复杂度: O(n) 需要额外的空间来存储每层的节点，累加为n
*/
func levelOrder429(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret := make([][]int, 0)
	nodeList := []*Node{root}
	for len(nodeList) > 0 {
		tmp := make([]int, 0)
		tmpNodeList := make([]*Node, 0)
		for _, v := range nodeList {
			tmp = append(tmp, v.Val)
			if len(v.Children) > 0 {
				tmpNodeList = append(tmpNodeList, v.Children...)
			}
		}
		ret = append(ret, tmp)
		nodeList = tmpNodeList
	}
	return ret
}

/*
思路二：递归，按照每层节点进行递归，遍历当前层节点，然后得出下一层节点，递归调用遍历函数
时间复杂度: O(n)
空间复杂度:O(n)
*/

func levelOrder429_2(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ret := make([][]int, 0)
	levelOrder_([]*Node{root}, &ret)
	return ret
}

func levelOrder_(nodes []*Node, ret *[][]int) {
	if len(nodes) < 1 {
		return
	}
	nextNodes := make([]*Node, 0)
	tmp := make([]int, 0)
	for _, v := range nodes {
		if len(v.Children) > 0 {
			nextNodes = append(nextNodes, v.Children...)
		}
		tmp = append(tmp, v.Val)
	}
	*ret = append(*ret, tmp)
	levelOrder_(nextNodes, ret)
}

func main() {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node1.Children = []*Node{node2, node3, node4}
	node3.Children = []*Node{node5, node6}
	fmt.Println(levelOrder429_2(node1))
}
