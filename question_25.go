package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head
	}
	ht := &ListNode{Next: head}
	prev := ht
	for {
		ok, last := hasKNode(prev, k)
		if !ok {
			break
		}
		next := last.Next
		newFirst, newLast := reverse25(prev.Next, last)
		prev.Next = newFirst
		newLast.Next = next

		prev = newLast
	}

	return ht.Next
}

func hasKNode(prev *ListNode, k int) (bool, *ListNode) {
	p := prev
	for i := 0; i < k; i++ {
		if p.Next != nil {
			p = p.Next
		} else {
			return false, nil
		}
	}
	return true, p
}

func reverse25(first, last *ListNode) (*ListNode, *ListNode) {
	p, q := first, first.Next
	for p != last {
		t := q.Next
		q.Next = p
		p = q
		q = t
	}
	return last, first
}

//func main() {
//	n1 := &ListNode{Val: 1}
//	n2 := &ListNode{Val: 2}
//	n3 := &ListNode{Val: 3}
//	n4 := &ListNode{Val: 4}
//	n5 := &ListNode{Val: 5}
//	n1.Next = n2
//	n2.Next = n3
//	n3.Next = n4
//	n4.Next = n5
//	l := n1
//
//	head := reverseKGroup(l, 2)
//	for ; head != nil; head = head.Next {
//		fmt.Printf("%d->", head.Val)
//	}
//}
