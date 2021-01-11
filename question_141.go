package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    quick := head
    slow := head
    for {
        slow = slow.Next
        if slow == nil {
            return false
        }
        quick = quick.Next
        if quick == nil {
            return false
        }
        quick = quick.Next
        if quick == nil {
            return false
        }
        if slow == quick {
            return true
        }
    }
}
