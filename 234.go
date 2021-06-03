package main

// 思路：找到中点，反转后半部分，(记得断开两个链表), 遍历对比是否相等
func isPalindromeL(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tail := reverse(slow.Next)
	slow.Next = nil
	for head != nil && tail != nil {
		if head.Val != tail.Val {
			return false
		}
		head = head.Next
		tail = tail.Next
	}
	return true
}

func r(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pre
		pre = head
		head = tmp
	}
	return pre
}
