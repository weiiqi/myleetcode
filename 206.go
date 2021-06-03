package main

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := new(ListNode)
	for head != nil {
		if ret.Next != nil {
			newNode := &ListNode{Val: head.Val}
			newNode.Next = ret.Next
			ret.Next = newNode
		} else {
			ret.Next = &ListNode{Val: head.Val}
		}
		head = head.Next
	}
	return ret.Next
}

func reverseList1(head *ListNode) *ListNode {
	prev := new(ListNode)
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}

/*
2021-05-28 链表反转
遍历方法:
这次思路完全正确，但是实现的时候有点乱，重点在
		head.Next = reverseHead
		reverseHead = head
时间复杂度 O(n)
空间复杂度 O(1)
*/

func reverseList20210528(head *ListNode) *ListNode {
	var reverseHead *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = reverseHead
		reverseHead = head
		head = tmp
	}
	return reverseHead
}

/*
2021-05-28 递归方法
n1 -> n2 -> n3 -> n4
1. head = n1, newhead = reverse(n2)
	2. head = n2, newhead =reverse(n3)
		3. head = n3, newhead = reverse(n4)
			4. head = n4, return n4
		3.1 head = n3, newhead = n4, head.next.next(n4.next) = head => n4.next = n3 , head.next = nil => n3.next = nil
		3.2 此时链表状态为 newhead = n4 -> n3 -> nil, head = n3
	2.1 head = n2, newhead = n4 -> n3 -> nil, head.next.next = head => n3.next = n2, head.next=nil => n2.next = nil
	2.2 此时链表状态为 n4 -> n3 -> n2 -> nil, head = n2
1.1 head = n1, head.next.next = head => n2.next = n1, head.next = nil => n1.next = nil
1.2 最终出循环 newhead = n4 -> n3 -> n2 -> n1 -> nil

时间复杂度为：O(n)
空间复杂度为：O(n) 堆栈深度为链表元素个数
*/

func reverseList20210528_1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList20210528_1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
