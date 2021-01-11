package leetcode

import (
    "fmt"
    "testing"
)

func TestMergeTwoLists(t *testing.T) {
    l1 := new(ListNode)
    l1.Val = 1
    l12 := new(ListNode)
    l12.Val = 2
    l1.Next = l12
    l13 := new(ListNode)
    l13.Val = 4
    l12.Next = l13

    l2 := new(ListNode)
    l2.Val = 1
    l22 := new(ListNode)
    l22.Val = 3
    l2.Next = l22
    l23 := new(ListNode)
    l23.Val = 4
    l22.Next = l23

    l := mergeTwoLists(l1, l2)
    for l != nil {
        fmt.Printf("%d->", l.Val)
        l = l.Next
    }
    t.Logf("\n")
}
