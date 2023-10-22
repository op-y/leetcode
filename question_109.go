package leetcode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	root := genTree(count)

	pos := head
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root.Left != nil {
			inorder(root.Left)
		}
		root.Val = pos.Val
		pos = pos.Next
		if root.Right != nil {
			inorder(root.Right)
		}
	}
	inorder(root)
	return root
}

func genTree(n int) *TreeNode {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return &TreeNode{}
	}
	if n == 2 {
		left := &TreeNode{}
		root := &TreeNode{Left: left}
		return root
	}
	if n == 3 {
		left := &TreeNode{}
		right := &TreeNode{}
		root := &TreeNode{Left: left, Right: right}
		return root
	}

	mid := n / 2
	nleft, nright := mid, n-mid-1
	root := &TreeNode{}
	left := genTree(nleft)
	right := genTree(nright)
	root.Left = left
	root.Right = right
	return root
}
