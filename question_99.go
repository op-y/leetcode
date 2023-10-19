package leetcode

func recoverTree(root *TreeNode)  {
	stack := []*TreeNode{}
	var x, y, prev *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left	
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil && root.Val < prev.Val {
			y = root
			if x == nil {
				x = prev
			} else {
				break
			}
		}
		prev = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
