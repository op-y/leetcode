package leetcode

func rightSideView(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }

    current := []*TreeNode{root}

    for {
        if len(current) == 0 {
            break
        }
        val := current[len(current)-1].Val
        result = append(result, val)

        next := []*TreeNode{}
        for _, node := range current {
            if node.Left != nil {
                next = append(next, node.Left)
            }
            if node.Right != nil {
                next = append(next, node.Right)
            }
        }
        current = next
    }

    return result
}
