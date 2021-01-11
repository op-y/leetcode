package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := minDepth(root.Left)
    rightDepth := minDepth(root.Right)
    if root.Left == nil {
        return 1 + rightDepth
    } else if root.Right == nil {
        return 1 + leftDepth
    } else {
        if leftDepth > rightDepth {
            return 1 + rightDepth
        } else {
            return 1 + leftDepth
        }
    }
}
