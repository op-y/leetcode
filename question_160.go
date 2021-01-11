package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }
    if headA == headB {
        return headA
    }

    pA := headA
    pB := headB
    lenA := 1
    lenB := 1
    for pA.Next != nil {
        lenA++
        pA = pA.Next
    }
    for pB.Next != nil {
        lenB++
        pB = pB.Next
    }
    if pA != pB {
        return nil
    }
    pA = headA
    pB = headB
    diff := 0
    if lenA >= lenB {
        diff = lenA - lenB
        for i:=0; i<diff; i++ {
            pA = pA.Next
        }
        for pA != pB {
            pA = pA.Next
            pB = pB.Next
        }
        return pA
    } else {
        diff = lenB - lenA
        for i:=0; i<diff; i++ {
            pB = pB.Next
        }
        for pA != pB {
            pA = pA.Next
            pB = pB.Next
        }
        return pA
    }
}
