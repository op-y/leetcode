package leetcode

import (
	"math"
)

func getMinimumDifference(root *TreeNode) int {
	mindiff := math.MaxInt

	var inorder func(root *TreeNode) (int, int)
	inorder = func(root *TreeNode) (int, int) {
		if root.Left == nil && root.Right == nil {
			return root.Val, root.Val
		}

		tmin, tmax := root.Val, root.Val
		if root.Left != nil {
			lmin, lmax := inorder(root.Left)
			mindiff = min(mindiff, root.Val-lmax)
			tmin = lmin
		}

		if root.Right != nil {
			rmin, rmax := inorder(root.Right)
			mindiff = min(mindiff, rmin-root.Val)
			tmax = rmax
		}

		return tmin, tmax
	}
	inorder(root)
	return mindiff
}
