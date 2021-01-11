package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    p1 := head
    p2 := head.Next
    for p2 != nil {
        p3 := p2.Next
        p2.Next = p1
        p1 = p2
        p2 = p3
    }
    head.Next = nil
    return p1
}
