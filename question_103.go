package leetcode

func zigzagLevelOrder(root *TreeNode) [][]int {
        result := [][]int{}
        if root == nil {
            return result
        }
        prev := []*TreeNode{root}
        reverse := false
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
            if reverse {
                reverseArray(level)
            }
            result = append(result, level)
            prev = next
            reverse = !reverse
        }
        return result
}

func reverseArray(level []int) {
    i:=0
    j:=len(level)-1
    for i<j {
        level[i], level[j] = level[j], level[i]
        i++
        j-- 
    }
}
