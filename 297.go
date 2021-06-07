package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
思路一： 通过前序遍历(根左右，较好确定前后左右关系)确定序列化的顺序, 根据前序规则对字符串进行反序列化，递归实现
时间复杂度: O(n)
空间复杂度: O(n)
*/

type Codec struct {
	codec []string
}

func Constructor() Codec {
	return Codec{}
}

func doSerialize(root *TreeNode, ret string) string {
	if root == nil {
		ret += "null,"
	} else {
		ret += strconv.Itoa(root.Val) + ","
		ret = doSerialize(root.Left, ret)
		ret = doSerialize(root.Right, ret)
	}
	return ret
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return doSerialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	l := strings.Split(data, ",")
	for _, v := range l {
		if v != "" {
			this.codec = append(this.codec, v)
		}
	}
	return this.deserialize_1()
}

func (this *Codec) deserialize_1() *TreeNode {
	if this.codec[0] == "null" {
		this.codec = this.codec[1:]
		return nil
	}
	// 前序遍历：根左右
	val, _ := strconv.Atoi(this.codec[0])
	this.codec = this.codec[1:]
	root := &TreeNode{Val: val, Left: this.deserialize_1(), Right: this.deserialize_1()}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
