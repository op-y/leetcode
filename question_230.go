package leetcode

// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	count := 0
	numk := 0
	found := false

	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if found {
			return
		}
		if node.Left != nil {
			inorder(node.Left)
		}

		count++
		if count == k {
			found = true
			numk = node.Val
		}
		
		if found {
			return
		}
		if node.Right != nil {
			inorder(node.Right)
		}
	}
	inorder(root)

	return numk
}


