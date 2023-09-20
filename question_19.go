package leetcode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil {
        return nil
    }

    pf := head
    ps := head
    move := 0
    for i:=0; i < n; i++ {
        if pf.Next != nil {
            pf = pf.Next
            move++
        } else {
            if move == n - 1 {
                break
            } else {
                // n大于sz的异常情况
                return nil
            }
        }
    }

    // 快指针移动到尾部
    for pf.Next != nil {
        pf = pf.Next
        ps = ps.Next
    }

    // 删除头节点的情况
    if move == n - 1 {
        if ps.Next == nil {
            return nil
        } else {
            ps.Val = ps.Next.Val
            ps.Next = ps.Next.Next
            return head
        }
    } else {
        ps.Next = ps.Next.Next
        return head
    }
}
