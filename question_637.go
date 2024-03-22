package leetcode

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		count, sum := 0, 0
		next := []*TreeNode{}

		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
			count++
			sum += node.Val
		}
		result = append(result, float64(sum)/float64(count))
		q = next	
	}

	return result
}


