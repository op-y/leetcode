package leetcode

type BSTIterator struct {
	root *TreeNode
	pos  *TreeNode
}

func Constructor173(root *TreeNode) BSTIterator {
	iter := BSTIterator{root: root}
	if root != nil {
		var cur1, cur2 *TreeNode
		cur1 = root
		cur2 = nil
		for cur1 != nil {
			cur2 = cur1.Left
			if cur2 != nil {
				// 当前节点的前驱节点
				for cur2.Right != nil && cur2.Right != cur1 {
					cur2 = cur2.Right
				}
				if cur2.Right == nil {
					cur2.Right = cur1 // 建立前驱节点连接
					cur1 = cur1.Left  // 再处理左子树
					continue
				}
				// 创建过程中不解开前驱链
				// 创建过程中不处理值
			}
			cur1 = cur1.Right // 利用之前建立的前驱节点返回到根节点
		}

		iter.pos = root
	}

	return iter
}

func (this *BSTIterator) Next() int {
	if this.pos != nil {
		cur := this.pos
		for this.pos.Left != nil {
			this.pos = this.pos.Left
			cur.Left = nil
			cur = this.pos
		}
		ans := this.pos.Val
		this.pos = this.pos.Right
		return ans
	}
	return -1 // 已经完成遍历
}

func (this *BSTIterator) HasNext() bool {
	if this.pos == nil {
		return false
	}
	return true
}

//func main() {
//	n0 := &TreeNode{Val: 7}
//	n1 := &TreeNode{Val: 3}
//	n0.Left = n1
//	n2 := &TreeNode{Val: 15}
//	n0.Right = n2
//	n3 := &TreeNode{Val: 9}
//	n4 := &TreeNode{Val: 20}
//	n2.Left = n3
//	n2.Right = n4
//
//	iter := Constructor173(n0)
//	fmt.Println(iter.Next())
//	fmt.Println(iter.Next())
//	fmt.Println(iter.HasNext())
//	fmt.Println(iter.Next())
//	fmt.Println(iter.HasNext())
//	fmt.Println(iter.Next())
//	fmt.Println(iter.HasNext())
//	fmt.Println(iter.Next())
//	fmt.Println(iter.HasNext())
//}
//
