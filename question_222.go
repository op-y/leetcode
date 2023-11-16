package leetcode

import (
	"math"
)

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	p := root
	for p.Left != nil {
		p = p.Left
		level++
	}
	low, high := int(math.Pow(2, float64(level))), int(math.Pow(2, float64(level+1))-1)
	for low <= high {
		mid := (low + high) / 2
		if search222(root, mid, level) {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return low - 1
}

func search222(root *TreeNode, id int, level int) bool {
	p := root
	for i := 1; i <= level; i++ {
		mask := 1 << (level - i)
		bit := id & mask
		bit >>= (level - i)
		if bit == 0 {
			p = p.Left
		} else {
			p = p.Right
		}
	}
	return p != nil
}

//func main() {
//	n1 := &TreeNode{Val: 1}
//	n2 := &TreeNode{Val: 1}
//	n3 := &TreeNode{Val: 1}
//	n4 := &TreeNode{Val: 1}
//	//n5 := &TreeNode{Val: 1}
//	//n6 := &TreeNode{Val: 1}
//	n1.Left = n2
//	n1.Right = n3
//	n2.Left = n4
//	//n2.Right = n5
//	//n3.Left = n6
//	r := countNodes(n1)
//	fmt.Println(r)
//}
