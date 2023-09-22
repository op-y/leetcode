package leetcode

func mergeTwoLists1801(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, lm, lm2, ls *ListNode
	if l1.Val < l2.Val {
		head = l1
		lm = l1
		lm2 = l1
		ls = l2
	} else {
		head = l2
		lm = l2
		lm2 = l2
		ls = l1
	}

	for {
		if lm.Val < ls.Val {
			if lm.Next == nil { // 主链已经到末尾了拼接次链退出
				lm.Next = ls
				break
			} else { // 主链后移
				if lm2 != lm {
					lm2 = lm2.Next
				}
				lm = lm.Next
			}
		} else {
			lnp := ls
			ls = ls.Next
			if lm != lm2 {
				lnp.Next = lm
				lm2.Next = lnp
				lm2 = lnp
			} else {
				lnp.Next = lm.Next
				lm2.Next = lnp
				lm = lnp
			}
			if ls == nil { // 次链已经到达末尾了直接退出
				break
			}
		}
	}

	return head
}
