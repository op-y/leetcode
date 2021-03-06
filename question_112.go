package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        if root.Val == sum {
            return true
        } else {
            return false
        }
    }

    if root.Left != nil {
        if hasPathSum(root.Left, sum - root.Val) {
            return true
        }
    }
    if root.Right != nil {
        if hasPathSum(root.Right, sum - root.Val) {
            return true
        }
    }

    return false
}
