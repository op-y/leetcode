package leetcode

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}	
	if root == nil {
		return ans
	}
	var traversal func(*TreeNode)
	traversal = func(root *TreeNode) {
		if root.Left != nil {
			traversal(root.Left)
		}
		ans = append(ans, root.Val)
		if root.Right != nil {
			traversal(root.Right)
		}
	}
	traversal(root)
	return ans
}

