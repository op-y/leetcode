package leetcode

func deleteDuplicates83(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	h := &ListNode{Next: head}
	p, i := h, h
	elem := i.Next.Val
	count := 0
	for i.Next != nil {
		i = i.Next
		if i.Val == elem {
			count++
		} else {
			if count == 1 {
				p = p.Next
			} else {
				p.Next = i
			}
			elem = i.Val
			count = 1
		}
	}
	if count == 1 {
		p = p.Next
	} else {
		p.Next = nil
	}

	return h.Next
}
