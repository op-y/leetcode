package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }

    p1 := head
    for p1 != nil {
        if p1.Val == val {
            p1 = p1.Next
        } else {
            break
        }
    }
    if p1 == nil {
        return nil
    }

    phead := p1
    p2 := p1
    if p1.Next == nil {
        return p1
    } else {
        p2 = p1.Next
    }

    for p2 != nil {
        if p2.Val == val {
            p2 = p2.Next
            p1.Next = p2
        } else {
            p1 = p2
            p2 = p2.Next
        }
    }
    return phead
}
