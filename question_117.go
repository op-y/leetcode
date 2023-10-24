package leetcode

func connect117(root *Node) *Node {
	if root == nil {
		return nil
	}
	preorder117(root)
	layer117(root)
	return root
}

func preorder117(root *Node) {
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Left != nil {
		step1(root.Left)
	}
	if root.Right != nil {
		step1(root.Right)
	}
	return
}

func layer117(root *Node) {
	ml, pos := root, root
	for ml != nil {
		pos = ml
		for pos != nil {
			var l, r *Node

			for l == nil && pos != nil {
				if pos.Right != nil && pos.Right.Next == nil {
					l = pos.Right

				} else if pos.Left != nil && pos.Left.Next == nil {
					l = pos.Left
				}
				// 找到左节点后需要后移 因为左右节点不可能在同一个节点下
				pos = pos.Next
			}

			for r == nil && pos != nil {
				if pos.Left == nil && pos.Right == nil {
					pos = pos.Next
					continue
				}
				// 找到右节点后 原地不动 下一个左节点可能在这里
				if pos.Left != nil {
					r = pos.Left
				} else if pos.Right != nil {
					r = pos.Right
				}
			}

			if l != nil && r != nil {
				l.Next = r
				l, r = nil, nil
			}
		}
		ml = mostLeft(ml) // 下一层
	}
}

func mostLeft(root *Node) *Node {
	for root != nil {
		if root.Left != nil {
			return root.Left
		}
		if root.Right != nil {
			return root.Right
		}
		root = root.Next
	}
	return nil
}
