package leetcode

import (
    "fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }
    return getPath(p, root)
}

func getPath(p string, node *TreeNode) []string {
    result := make([]string, 0)
    var path string
    if p == "" {
        path = fmt.Sprintf("%d", node.Val)
    } else {
        path = fmt.Sprintf("%s->%d", p, node.Val)
    }

    if node.Left == nil && node.Right == nil {
        result = append(result, path)
        return result
    }

    if node.Left != nil {
        left := getPath(path, node.Left)
        result = append(result, left...)
    }
    if node.Right != nil {
        right := getPath(path, node.Right)
        result = append(result, right...)
    }
    return result
}
