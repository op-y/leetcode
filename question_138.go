package leetcode

func copyRandomList(head *RandomNode) *RandomNode {
	if head == nil {
		return nil
	}

	// 复制节点
	p := head
	next := p.Next
	for p != nil {
		n := &RandomNode{Val: p.Val}
		p.Next = n
		n.Next = next
		p = next
		if p != nil {
			next = p.Next
		} else {
			next = nil
		}
	}

	// 处理random
	p = head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}

    newHead := head.Next
	oldp, newp := head, newHead
    var oldNext, newNext *RandomNode
	for oldp != nil && newp != nil {
		oldNext = oldp.Next.Next
		if newp.Next != nil {
			newNext = newp.Next.Next
		} else {
			newNext = nil
		}

		// 隔一个位置取下个节点
		oldp.Next = oldNext
		newp.Next = newNext

		// 后移
		oldp = oldNext
		newp = newNext
	}

	return newHead
}
