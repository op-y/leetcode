leetcode

func isValidBST(root *TreeNode) bool {
    ok, _, _ := check(root)
    return ok
}

func check(root *TreeNode) (bool, int, int) {
    min, max := root.Val, root.Val
    if root.Left == nil && root.Right == nil  {
        return true, min, max
    }
    if root.Left != nil {
        ok, lmin, lmax := check(root.Left)
        if !ok || root.Val <= lmax {
            return false, min, max
        }
        min = lmin
    }
    if root.Right != nil {
        ok, rmin, rmax := check(root.Right)
        if !ok || root.Val >= rmin {
            return false, min, max
        }
        max = rmax
    }

    return true, min, max
}
