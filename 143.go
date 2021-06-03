package main

// 找到中点，断开，翻转后半部分，合并两部分
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	mid := findMid(head)
	tail := reverse(mid.Next)
	mid.Next = nil
	head = merge(head, tail)
	return
}

func findMid(head *ListNode) *ListNode {
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

func merge(l1, l2 *ListNode) *ListNode {
	tmp := new(ListNode)
	head := tmp
	flag := true
	for l1 != nil && l2 != nil {
		if flag {
			head.Next = l1
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		flag = !flag
		head = head.Next
	}
	for l1 != nil {
		head.Next = l1
		l1 = l1.Next
		head = head.Next
	}
	for l2 != nil {
		head.Next = l2
		l2 = l2.Next
		head = head.Next
	}
	return tmp.Next
}
