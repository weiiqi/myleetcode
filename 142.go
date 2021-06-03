package main

// 通过hash，如果有>2 的就是环入口起点
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := make(map[*ListNode]int)
	for head != nil && head.Next != nil {
		_, ok := tmp[head.Next]
		if ok {
			return head.Next
		} else {
			tmp[head]++
		}
		head = head.Next
	}
	return nil
}

// 快慢指针, 快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇带你即为入环点
func detectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head
	for fast != nil && fast.Next != nil {
		if fast == slow {
			slow = head
			fast = fast.Next
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}

/*
2021-05-28
思路一：通过hash table, 存储链表，重复访问到的即是环开始的地方
*/
func detectCycle20210528(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	hasht := make(map[*ListNode]struct{})
	for head != nil {
		if _, ok := hasht[head]; ok {
			return head
		}
		hasht[head] = struct{}{}
		head = head.Next
	}
	return nil
}

/*
思路二： 快慢指针
// 错误思维：相遇点不一定是环开始的地方
链表：进入环要走X的举例，起点与相遇点的举例为Y， 相遇点与起点的具体为Z，
慢指针走过的路程： x+y
快指针走过的路程: x + y + z + y
fast = 2 * slow
故，x = z
*/

func detectCycle20210528_1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			slow = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return slow
		}
	}
	return nil
}
