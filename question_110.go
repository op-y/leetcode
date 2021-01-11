package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    _, isBalance := getHeight(root)
    return isBalance
}

func getHeight(root *TreeNode) (int, bool) {
    if root == nil {
        return 0, true
    }

    leftHeight, leftBalance := getHeight(root.Left)

    rightHeight, rightBalance := getHeight(root.Right)

    if !leftBalance || !rightBalance {
        return -1, false
    }

    if leftHeight - rightHeight > 1 {
        return -1, false
    }

    if rightHeight - leftHeight > 1 {
        return -1, false
    }

    height := 0
    if leftHeight >= rightHeight {
        height = leftHeight + 1
    } else {
        height = rightHeight + 1
    }
    return height, true
}
