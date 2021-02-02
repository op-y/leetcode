package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var head, cursor *ListNode
    carrier := false

    for l1 != nil && l2 != nil {
        value := l1.Val + l2.Val
        if carrier {
            value += 1
            carrier = false
        }
        if value >= 10 {
            carrier = true
            value -= 10
        }
        if head == nil {
            node := new(ListNode)
            node.Val = value
            head = node
            cursor = node
        } else {
            node := new(ListNode)
            node.Val = value
            cursor.Next = node
            cursor = node
        }
        l1 = l1.Next
        l2 = l2.Next
    }

    for l1 != nil {
        value := l1.Val
        if carrier {
            value += 1
            carrier = false
        }
        if value >= 10 {
            carrier = true
            value -= 10
        }
        node := new(ListNode)
        node.Val = value
        cursor.Next = node
        cursor = node
        l1 = l1.Next
    }

    for l2 != nil {
        value := l2.Val
        if carrier {
            value += 1
            carrier = false
        }
        if value >= 10 {
            carrier = true
            value -= 10
        }
        node := new(ListNode)
        node.Val = value
        cursor.Next = node
        cursor = node
        l2 = l2.Next
    }

    if carrier {
        node := new(ListNode)
        node.Val = 1
        cursor.Next = node
        cursor = node
    }

    return head
}
