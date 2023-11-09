package leetcode

//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	nh, _ := partition(head)
//	return nh
//}
//
//func partition(head *ListNode) (*ListNode, *ListNode) {
//	if head == nil {
//		return nil, nil
//	}
//	if head.Next == nil {
//		return head, head
//	}
//
//	var hlt, tlt, hgt, tgt *ListNode
//	hlt, hgt = &ListNode{}, &ListNode{}
//	tlt = hlt
//	tgt = hgt
//
//	// 选头节点为标志
//	flag := head
//	// 从第二个节点开始处理
//	pos := head.Next
//	// 头节点从链上断开
//	flag.Next = nil
//
//	for pos != nil {
//		// 小于头节点的放到小链
//		// 大于头节点的放到大链
//		if pos.Val < flag.Val {
//			tlt.Next = pos
//			tlt = pos
//		} else {
//			tgt.Next = pos
//			tgt = pos
//		}
//		tmp := pos
//		pos = pos.Next
//		tmp.Next = nil // 已经处理的节点从链上断开
//	}
//	// 递归处理大小链
//	hlt = hlt.Next
//	hgt = hgt.Next
//	nhlt, ntlt := partition(hlt)
//	nhgt, ntgt := partition(hgt)
//
//	// 合并大小链和标志节点
//	var nh, nt *ListNode
//	if nhlt != nil && ntlt != nil {
//		nh = nhlt
//		nt = ntlt
//		nt.Next = flag
//		nt = flag
//	} else {
//		nh, nt = flag, flag
//	}
//	if nhgt != nil && ntgt != nil {
//		nt.Next = nhgt
//		nt = ntgt
//	}
//
//	return nh, nt
//}
