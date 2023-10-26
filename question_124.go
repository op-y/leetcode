package leetcode

func maxPathSum(root *TreeNode) int {
	found := false
	max := 0
	var postorder func(*TreeNode) int
	postorder = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			if !found || root.Val > max {
				max = root.Val
				found = true
			}
			return root.Val
		}

		lmax, rmax := 0, 0
		if root.Left != nil {
			lmax = postorder(root.Left)
		}
		if root.Right != nil {
			rmax = postorder(root.Right)
		}

		curMax := max124(root.Val, max124(root.Val+lmax+rmax, max124(root.Val+lmax, root.Val+rmax)))
		if !found || curMax > max {
			max = curMax
			found = true
		}
		retMax := max124(root.Val, max124(root.Val+lmax, root.Val+rmax))
		return retMax
	}
	postorder(root)
	return max
}

func max124(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//func main() {
//	root := &TreeNode{Val: -3}
//
//	n1 := &TreeNode{Val: 2}
//	root.Left = n1
//
//	n2 := &TreeNode{Val: 3}
//	n3 := &TreeNode{Val: -6}
//	n1.Left = n2
//	n1.Right = n3
//
//	n4 := &TreeNode{Val: 2}
//	n3.Left = n4
//
//	max := maxPathSum(root)
//	fmt.Printf("%v\n", max)
//}
