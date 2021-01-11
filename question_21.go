package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
   if l1 == nil && l2 == nil {
        return nil
   }
   header := new(ListNode)
   current := header
   for l1 != nil && l2 != nil {
        next := new(ListNode)
        if l1.Val <= l2.Val {
            next.Val = l1.Val
            l1 = l1.Next
        } else {
            next.Val = l2.Val
            l2 = l2.Next
        }
        current.Next = next
        current = next
   }
   if l1 == nil {
       for l2 != nil {
           next := new(ListNode)
           next.Val = l2.Val
           l2 = l2.Next
           current.Next = next
           current = next
       }
   }
   if l2 == nil {
       for l1 != nil {
           next := new(ListNode)
           next.Val = l1.Val
           l1 = l1.Next
           current.Next = next
           current = next
       }
   }
   return header.Next
}
