package main

func sortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func mergeSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil

	left := mergeSort(head)
	right := mergeSort(tail)
	return mergeTwoLists(left, right)
}
