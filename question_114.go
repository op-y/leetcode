package leetcode

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	genList(root)
}


func genList(root *TreeNode) (*TreeNode, *TreeNode) {
	if root.Left == nil && root.Right == nil {
		return root, root
	}
	var lhead, ltail, rhead, rtail *TreeNode
	if root.Left != nil {
		lhead, ltail = genList(root.Left)
	}
	if root.Right != nil {
		rhead, rtail = genList(root.Right)
	}

	if lhead != nil && rhead != nil {
		root.Right = lhead
		root.Left = nil	
		ltail.Right = rhead
		return root, rtail
	} else if lhead != nil {
		root.Right = lhead
		root.Left = nil
		return root, ltail
	} else if rhead != nil {
		root.Right = rhead
		root.Left = nil
		return root, rtail
	} else {
		return root, root
	}	
}
