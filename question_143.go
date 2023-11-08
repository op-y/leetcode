package leetcode

func reorderList(head *ListNode) {
	// 没有节点
	if head == nil {
		return
	}
	// 只有一个节点
	if head.Next == nil {
		return
	}
	// 只有两个节点
	if head.Next.Next == nil {
		return
	}

	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil { // fast 走得快
		if slow.Next != nil {
			slow = slow.Next
		}
		if fast.Next != nil {
			fast = fast.Next
		}
		if fast.Next != nil {
			fast = fast.Next
		}
	}

	first := head
	second := slow.Next
	slow.Next = nil

	second, _ = reverse143(second)
	for second != nil {
		tmp := second
		second = second.Next

		tmp.Next = first.Next
		first.Next = tmp
		first = tmp.Next
	}
	return
}

func reverse143(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}
	if head.Next == nil {
		return head, head
	}
	h, t := reverse143(head.Next)
	t.Next = head
	t = head
	t.Next = nil
	return h, t
}

//func main() {
//	n0 := &ListNode{Val: 1}
//	n1 := &ListNode{Val: 2}
//	n2 := &ListNode{Val: 3}
//	n3 := &ListNode{Val: 4}
//	n0.Next = n1
//	n1.Next = n2
//	n2.Next = n3
//	n3.Next = n1
//	reorderList(n0)
//}
