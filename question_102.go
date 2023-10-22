package leetcode

func levelOrder(root *TreeNode) [][]int {
		result := [][]int{}
		if root == nil {
			return result
		}
		prev := []*TreeNode{root}
		for len(prev) != 0 {
			level := []int{}
			next := []*TreeNode{}
			for _, node := range prev {
				level = append(level, node.Val)
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
			result = append(result, level)
			prev = next
		}
		return result
}
