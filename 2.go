package main


type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    delt,total := 0,0
    ret := &ListNode{}
    head := ret

    for (l1 != nil ||l2 != nil) || (delt != 0){
        head.Next = new(ListNode)
        head = head.Next
        total = 0
        if l1 != nil{
            total += l1.Val
            l1 = l1.Next
        }
        if l2 != nil{
            total += l2.Val
            l2 = l2.Next
        }
        total += delt
        if total >= 10{
            delt = 1
            total = total -10
        }else{
            delt = 0
        }
        head.Val = total
    }
    return ret.Next
}