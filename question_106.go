package leetcode

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}
	if len(postorder) == 1 && len(inorder) == 1 {
		root := &TreeNode{Val: postorder[0]}
		return root
	}

	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var left, right *TreeNode

	idx := getPos106(inorder, postorder[len(postorder)-1])
	leftLen := idx
	rightLen := len(inorder) - idx - 1
	if leftLen > 0 {
		leftPostorder := postorder[0:idx]
		leftInorder := inorder[0:idx]
		left = buildTree(leftInorder, leftPostorder)
	}
	if rightLen > 0 {
		rightPostorder := postorder[idx : len(postorder)-1]
		rightInorder := inorder[idx+1:]
		right = buildTree(rightInorder, rightPostorder)
	}
	root.Left = left
	root.Right = right
	return root
}

func getPos106(inorder []int, val int) int {
	for idx, v := range inorder {
		if v == val {
			return idx
		}
	}
	return -1
}
