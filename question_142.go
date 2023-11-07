package leetcode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}

	var slow, quick *ListNode
	if head.Next != nil {
		slow = head.Next
	} else {
		return nil
	}
	if slow.Next != nil {
		quick = slow.Next
	} else {
		return nil
	}

	for slow != quick {
		if slow.Next != nil {
			slow = slow.Next
		} else {
			return nil
		}

		if quick.Next != nil && quick.Next.Next != nil {
			quick = quick.Next.Next
		} else {
			return nil
		}
	}

	// 特殊情况
	if head == slow {
		return head
	}

	// 环的长度
	p := slow.Next
	cycleLen := 1
	for p != slow {
		p = p.Next
		cycleLen++
	}

	// 判断head是否在环内
	pos := head
	for i := 0; i < cycleLen; i++ {
		pos = pos.Next
	}
	if pos == head {
		return head
	}

	pos = head
	for pos != slow {
		pos = pos.Next
		slow = slow.Next
	}
	return pos
}

//func main() {
//	//n0 := &ListNode{Val: -1}
//	//n1 := &ListNode{Val: -7}
//	//n2 := &ListNode{Val: 7}
//	//n3 := &ListNode{Val: 4}
//	//n4 := &ListNode{Val: 19}
//	//n5 := &ListNode{Val: 6}
//	//n6 := &ListNode{Val: -9}
//	//n7 := &ListNode{Val: -5}
//	//n8 := &ListNode{Val: -2}
//	//n9 := &ListNode{Val: -5}
//	//n0.Next = n1
//	//n1.Next = n2
//	//n2.Next = n3
//	//n3.Next = n4
//	//n4.Next = n5
//	//n5.Next = n6
//	//n6.Next = n7
//	//n7.Next = n8
//	//n8.Next = n9
//	//n9.Next = n6
//
//	n0 := &ListNode{Val: 3}
//	n1 := &ListNode{Val: 2}
//	n2 := &ListNode{Val: 0}
//	n3 := &ListNode{Val: -4}
//	n0.Next = n1
//	n1.Next = n2
//	n2.Next = n3
//	n3.Next = n1
//	detectCycle(n0)
//}
