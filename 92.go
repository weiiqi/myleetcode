package main

// 思路。先遍历到M处，反转，再拼接后续
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil{
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	head = dummy

	var prev *ListNode
	i:=0
	for i<m{
		prev = head
		head = head.Next
		i++
	}
	j := i
	mid := head
	next := new(ListNode)
	for head != nil && j<=n{
		tmp := head.Next
		head.Next = next
		next = head
		head = tmp
		j++
	}
	prev.Next = next
	mid.Next = head
	return dummy.Next
}