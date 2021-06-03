package main

import "fmt"

/*
这题其实逻辑不复杂，但是有点绕，头指针的引用，循环过程中指针的变化指向，都很重要
总的思路可以简单描述为：要把大象装冰箱总共分三步
1. 创建一个头指针，持用反转后链表头
2. 按照长度找出需要反转的链表的头部和尾部
3. 反转链表，将反转后的链表和头指针联合起来返回
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	newHead := new(ListNode)
	newHead.Next = head
	tmp := newHead
	for head != nil {
		i := 0
		var endStep *ListNode
		startNode := head
		endStep = tmp
		for i = 0; i < k; i++ {
			endStep = endStep.Next
			if endStep == nil {
				return newHead.Next
			}
		}
		startNode, endStep = reverseLink(startNode, endStep)
		tmp.Next = startNode
		tmp = endStep
		head = endStep.Next
	}

	return newHead.Next
}

func main() {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = nil
	fmt.Println(reverseKGroup(node1, 3))
}

/*
区别于链表反转
链表反转最后一个为空指针，而这次需要先指向尾部指针，再继续反转其他节点
*/
func reverseLink(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	p := head
	for pre != tail {
		tmp := p.Next
		p.Next = pre
		pre = p
		p = tmp
	}
	return tail, head
}
