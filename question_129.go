package leetcode

func sumNumbers(root *TreeNode) int {
	sum := 0
	var preorder func(*TreeNode, int)
	preorder = func(root *TreeNode, s int) {
		if root.Left == nil && root.Right == nil {
			num := s + root.Val
			sum += num
		}
		if root.Left != nil {
			preorder(root.Left, (s + root.Val) * 10)
		}
		if root.Right != nil {
			preorder(root.Right, (s + root.Val) * 10)
		}
	}
	preorder(root, 0)	
	return sum
}

