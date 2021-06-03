package main

func deleteDuplicates2(head *ListNode) *ListNode {
	curr := new(ListNode)
	curr.Next = head
	head = curr
	for head.Next != nil && head.Next.Next != nil{
		rmVal := 0
		if head.Next.Val == head.Next.Next.Val{
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal{
				head.Next = head.Next.Next
			}
		}else{
			head = head.Next
		}
	}
	return curr.Next
}


//TODO 82 删除重复元素
func deleteDuplicates3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	first := new(ListNode)
	first.Next = head
	tmp := first
	second := first.Next
	flag := false
	for second.Next !=  nil{
		if second.Val == second.Next.Val{
			flag = true
			second = second.Next
		}else{
			if flag{
				second = second.Next
				first.Next=second
				flag = false
			}else{
				first = first.Next
				second = second.Next
			}
		}
	}
	if flag{
		first.Next = second.Next
	}
	return tmp.Next
}
