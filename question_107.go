package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    return getLayerValue(root)
}

func getLayerValue(root *TreeNode) [][]int {
    layer := []*TreeNode{root}

    result := make([][]int, 0)

    for len(layer) > 0 {
        data := make([]int, 0)
        children := make([]*TreeNode, 0)
        for _, node := range layer {
            data = append(data, node.Val)
            if node.Left != nil {
                children = append(children, node.Left)
            }
            if node.Right != nil {
                children = append(children, node.Right)
            }
        }
        result = append([][]int{data}, result...)
        layer = children
    }
    return result
}
