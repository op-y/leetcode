package main

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return head
    }

	odd, even := split(head)
	h := join(even, odd)
    return h
}

func split(head *ListNode) (*ListNode, *ListNode) {
	odd := head
	even := head.Next

	po := odd
	pe := even

	for po != nil && pe != nil {
		if po.Next != nil {
			po.Next = po.Next.Next
			po = po.Next
		}
		if pe.Next != nil {
			pe.Next = pe.Next.Next
			pe = pe.Next
		}
	}

	return odd, even
}

func join(first, second *ListNode) *ListNode {
	head := first
	
	for first != nil {
		p := first.Next
		first.Next = second
		first = p

		p = second.Next
		second.Next = first

		if second.Next == nil && p != nil {
			second.Next = p
		}
		second = p
	}

	return head
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 2}
	n3 := &ListNode{Val: 3}
	n4 := &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	l := n1


	head := swapPairs(l)
	for ; head != nil; head = head.Next {
		fmt.Printf("%d->", head.Val)
	}
}
