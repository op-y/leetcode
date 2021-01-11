package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    isLastLayer := false
    children := []*TreeNode{root}
    for !isLastLayer {
        nextLayer := make([]*TreeNode, 0)
        isLastLayer = true
        for _, node := range children {
            // 还有子节点就有可能继续往下循环
            if node != nil {
                if node.Left != nil || node.Right != nil {
                    isLastLayer = false
                }
                nextLayer = append(nextLayer, node.Left, node.Right)
            } else {
                nextLayer = append(nextLayer, nil, nil)
            }
        }
        if isLastLayer {
            return true
        }
        h := 0
        t := len(nextLayer) - 1
        for h < t {
            if nextLayer[h] == nil && nextLayer[t] != nil {
                return false
            }
            if nextLayer[h] != nil && nextLayer[t] == nil {
                return false
            }
            if nextLayer[h] == nil && nextLayer[t] == nil {
                h++
                t--
                continue
            }
            if nextLayer[h].Val != nextLayer[t].Val {
                return false
            } else {
                h++
                t--
            }
        }
        children = nextLayer
    }
    return true
}
