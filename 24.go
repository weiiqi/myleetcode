package main

/*
2021-05-28
思路一：通过循环遍历链表，反转就近的两个，思路大概正确，但是实际写代码的时候存在困难，害，加油吧！
下面是参考了别人的写法，手写了一遍

时间复杂度为O(n)
空间复杂度为O(1)
*/

func swapPairs(head *ListNode) *ListNode {
	newHead := new(ListNode)
	newHead.Next = head
	tmp := newHead
	for tmp.Next != nil && tmp.Next.Next != nil {
		node1 := tmp.Next
		node2 := tmp.Next.Next

		// exchange node1, node2
		tmp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		tmp = node1
	}
	return newHead.Next
}

/*
方法二： 在官网还学习到另一种清晰的思路，但是性能不是很好
递归:
1. 选前两个节点交换
2. 递归反转剩下的链表
3. 最上层返回链表

时间复杂度O(n)
空间复杂度O(n)
*/
func swapPairs20210528_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairs20210528_1(newHead.Next) //
	newHead.Next = head
	return newHead
}
