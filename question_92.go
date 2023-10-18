package leetcode

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	nleft, nright := dummy, dummy
	for i := 0; i < left; i++ {
		nleft = nleft.Next
	}
	for i := 0; i < right; i++ {
		nright = nright.Next
	}

	// 记录反转区间前后节点
	prev := dummy
	for prev.Next != nleft {
		prev = prev.Next
	}
	next := nright.Next

	var reverse func(*ListNode, *ListNode) (*ListNode, *ListNode)
	reverse = func(h, t *ListNode) (*ListNode, *ListNode) {
		if h == t {
			return h, t
		}
		prev := h
		newh, newt := reverse(h.Next, t)
		newt.Next = prev
		prev.Next = nil
		return newh, prev
	}
	newh, newt := reverse(nleft, nright)
	prev.Next = newh
	newt.Next = next

	return dummy.Next
}
