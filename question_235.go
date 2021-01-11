package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    common := root
    for common != nil {
        if common.Val > p.Val && common.Val < q.Val {
            return common
        }
        if common.Val < p.Val && common.Val > q.Val {
            return common
        }

        if common.Val == p.Val && common.Val != q.Val {
            return common
        }

        if common.Val != p.Val && common.Val == q.Val {
            return common
        }

        if common.Val > p.Val && common.Val > q.Val {
            common = common.Left
            continue
        }

        if common.Val < p.Val && common.Val < q.Val {
            common = common.Right
            continue
        }
    }
    return nil
}
