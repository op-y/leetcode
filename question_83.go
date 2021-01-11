package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    position := head
    cursor := head.Next
    for cursor != nil {
        if cursor.Val == position.Val {
            cursor = cursor.Next
        } else {
            position.Next = cursor
            position = cursor
            cursor = cursor.Next
        }
    }
    position.Next = nil
    return head
}
