package leetcode

func pathSum113(root *TreeNode, targetSum int) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	stack := []int{}
	count := 0	
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		count += root.Val
		stack = append(stack, root.Val)
		if root.Left == nil && root.Right == nil {
			if count == targetSum {
				ans := make([]int, len(stack))
				copy(ans, stack)
				result = append(result, ans)
			}
		}
		if root.Left != nil {
			preorder(root.Left)
		}
		if root.Right != nil {
			preorder(root.Right)
		}
		count -= root.Val
		stack = stack[0:len(stack)-1]
	}
	preorder(root)
	return result
}
