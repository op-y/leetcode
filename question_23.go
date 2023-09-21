package main

import (
	"fmt"
	"os"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    if isAllEmpty(lists) {
        return nil
    }
    var h *ListNode = nil
    var p *ListNode = nil

    for {
        idx := pickMin(lists)
		fmt.Printf("idx %d\n", idx)
        if idx == -1 {
            break
        } else {
            if h == nil {
                h = lists[idx]
                p = lists[idx]
            } else {
				p.Next = lists[idx]
				p = p.Next
			}
            lists[idx] = lists[idx].Next
        }
    }
    return h
}

func pickMin(lists []*ListNode) int {
    idx := -1
    for i, nl := range lists {
        if nl == nil {
            continue
        }
        if idx == -1 {
            idx = i
            continue
        }
        if nl.Val < lists[idx].Val {
            idx = i
        }
    }
    return idx
}

func isAllEmpty(lists []*ListNode) bool {
    empty := 0
    for _, nl := range lists {
        if nl == nil {
            empty += 1
        }
    }
    return empty == len(lists)
}

func main() {
	n1 := &ListNode{Val: 1}
	n2 := &ListNode{Val: 4}
	n3 := &ListNode{Val: 5}
	n1.Next = n2
	n2.Next = n3
	l1 := n1

	n1 = &ListNode{Val: 1}
	n2 = &ListNode{Val: 3}
	n3 = &ListNode{Val: 4}
	n1.Next = n2
	n2.Next = n3
	l2 := n1

	n1 = &ListNode{Val: 2}
	n2 = &ListNode{Val: 6}
	n1.Next = n2
	l3 := n1

	lists := []*ListNode{l1, l2, l3}

	header := mergeKLists(lists)
	if header == nil {
		fmt.Printf("header %v: []", header)		
		os.Exit(0)
	}
	fmt.Printf("header %v\n", header)
	for ; header != nil; header = header.Next {
		fmt.Printf("%d->", header.Val)
	}	
}
