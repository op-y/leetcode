package leetcode

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	step1(root)
	step2(root)
	return root
}

func step1(root *Node) {
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


func step2(root *Node) {
	// 因为是完美二叉树这么操作才没有问题 否则得逐层操作见117
	if root.Next != nil {
		var l, r *Node
		if root.Right != nil {
			l = root.Right
		} else if root.Left != nil {
			l = root.Left
		}
		if root.Next.Left != nil {
			r = root.Next.Left
		} else if root.Next.Right != nil {
			r = root.Next.Right
		}
		if l != nil && r != nil {
			l.Next = r
		}
	}
	if root.Left != nil {
		step2(root.Left)
	}
	if root.Right != nil {
		step2(root.Right)
	}
}
