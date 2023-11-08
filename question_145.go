package leetcode

// morris 算法
func postorderTraversal145(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	dummy := &TreeNode{Left: root} // 创建一个dummy节点让原来的树完全处于左子树种
	var cur1, cur2 *TreeNode
	cur1 = dummy
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
			} else {
				cur2.Right = nil // 删除建立的前驱节点连接
				nums := reverseRightLine(cur1.Left)
				ans = append(ans, nums...)
			}
		}
		cur1 = cur1.Right // 利用之前建立的前驱节点返回
	}
	return ans
}

func reverseRightLine(node *TreeNode) []int {
	nums := []int{}
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Right
	}
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}

//func main() {
//	n1 := &TreeNode{Val: 1}
//	n2 := &TreeNode{Val: 2}
//	n3 := &TreeNode{Val: 3}
//	n1.Right = n2
//	n2.Left = n3
//	res := postorderTraversal145(n1)
//	fmt.Print("%v\n", res)
//}
