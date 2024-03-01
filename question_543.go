package leetcode

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0

	var preorder func(root *TreeNode) (int)
	preorder = func(root *TreeNode) (int) {
		lmax, rmax := 0, 0

		if root.Left == nil && root.Right == nil {
			return 0
		}

		if root.Left != nil {
			lmax = preorder(root.Left)
			lmax++
		}
		if root.Right != nil {
			rmax = preorder(root.Right)
			rmax++
		}
	
		if lmax + rmax > max {
			max = lmax + rmax
		}

		if lmax > rmax {
			return lmax
		} else {
			return rmax
		}
	}
	preorder(root)

	return max
}

