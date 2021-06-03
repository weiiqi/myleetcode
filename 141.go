package main

// 快慢指针, 快指针步长2， 慢指针步长1，快慢指针相等，则有环
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	tmp := make(map[*ListNode]int)
	for head != nil {
		_, ok := tmp[head]
		if ok {
			return true
		} else {
			tmp[head]++
		}
		head = head.Next
	}
	return false
}

/*
2021-05-28
通过hash表存储节点，遍历过程中如果再次遇到hash表内存在的节点即可视为存在环否则不存在
时间复杂度为O(n)
空间复杂度为O(n)
*/
func hasCycle20210528(head *ListNode) bool {
	hasht := make(map[*ListNode]bool)
	if head == nil {
		return false
	}
	for head.Next != nil {
		if _, ok := hasht[head]; ok {
			return true
		}
		hasht[head] = true
		head = head.Next
	}
	return false
}

/*
2021-05-28 学习快慢指针解法
通过两个指针，快指针是慢指针的速度的2倍，如果存在环，就会二者就会差速就会相遇，否则遍历完成就不会相遇
时间复杂度为O(n)
空间复杂度为O(1)
这种方式真的很妙！
*/

func hasCycle20210528_1(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func hasCycle20210528_2(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
