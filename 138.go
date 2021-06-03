package main

type Node1 struct {
	Val    int
	Next   *Node1
	Random *Node1
}

// TODO 复制带随机指针的链表
func copyRandomList(head *Node1) *Node1 {
	if head == nil {
		return head
	}
	// 复制节点，紧挨到到后面
	// 1->2->3  ==>  1->1'->2->2'->3->3'
	cur := head
	for cur != nil {
		clone := &Node1{Val: cur.Val, Next: cur.Next}
		temp := cur.Next
		cur.Next = clone
		cur = temp
	}
	// 处理random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	// 分离两个链表
	cur = head
	cloneHead := cur.Next
	for cur != nil && cur.Next != nil {
		temp := cur.Next
		cur.Next = cur.Next.Next
		cur = temp
	}
	// 原始链表头：head 1->2->3
	// 克隆的链表头：cloneHead 1'->2'->3'
	return cloneHead
}
