package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedArrayToBST(nums []int) *TreeNode {
    if nums == nil {
        return nil
    }

    length := len(nums)
    if length == 0 {
        return nil
    }

    return arrayToTree(0, length-1, nums)
}

func arrayToTree(low, high int, nums []int) *TreeNode {
    if low < 0 || high >= len(nums) {
        return nil
    }
    if low > high {
        return nil
    }

    mid := (low + high) / 2
    root := new(TreeNode)
    root.Val = nums[mid]
    root.Left = arrayToTree(low, mid-1, nums)
    root.Right = arrayToTree(mid+1, high, nums)

    return root
}
