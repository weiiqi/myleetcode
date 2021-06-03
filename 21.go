package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := new(ListNode)
	tmp := ret
	for l1 != nil && l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
		}
		if l2 != nil {
			n2 = l2.Val
		}
		if n1 < n2 {
			ret.Next = &ListNode{Val: n1}
			l1 = l1.Next
		} else {
			ret.Next = &ListNode{Val: n2}
			l2 = l2.Next
		}
		ret = ret.Next
	}
	if l1 != nil {
		ret.Next = l1
	}
	if l2 != nil {
		ret.Next = l2
	}
	return tmp.Next
}

/*
2021-06-01
思路一：
新建头指针merge两个链表
时间复杂度：O(n)
空间复杂度：O(1)
对比题解，属于最优解，开心，加油！
*/
func mergeTwoLists20210601(l1 *ListNode, l2 *ListNode) *ListNode {
	newHead := new(ListNode)
	tmp := newHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 != nil {
		tmp.Next = l1
	}
	if l2 != nil {
		tmp.Next = l2
	}
	return newHead.Next
}

/*
2021-06-01
思路二：递归,  看起来代码更加优雅
时间复杂度: O(n)
空间复杂度：O(n) 来自于递归堆栈
*/

func mergeTwoLinkList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLinkList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLinkList(l1, l2.Next)
		return l2
	}

}
