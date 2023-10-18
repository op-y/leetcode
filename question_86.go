package leetcode

func partition86(head *ListNode, x int) *ListNode {
	// 空链表和只有一个节点链表情况直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 头节点
	dummy := &ListNode{
		Next: head,
	}

	// 找到x之前第一个大于等于x的结点 有可能就是x本身
	prev_pos, pos := dummy, head
	for pos != nil && pos.Val < x {
		prev_pos = prev_pos.Next
		pos = pos.Next
	}
	if pos == nil {
		return head
	}

	// 从头开始遍历小于x的节点放到pos节点之前
	prev_flag, flag := pos, pos.Next
	for flag != nil {
		if flag.Val < x {
			prev_flag.Next = flag.Next
			flag.Next = pos
			prev_pos.Next = flag
			prev_pos = flag
			flag = prev_flag.Next
		} else {
			prev_flag = prev_flag.Next
			flag = flag.Next
		}
	}

	return dummy.Next
}
