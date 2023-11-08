package leetcode

// morris 算法
func preorderTraversal144(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    ans := []int{}
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
                ans = append(ans, cur1.Val) // 先处理根节点
                cur1 = cur1.Left // 再处理左子树
                continue
            } else {
                cur2.Right = nil // 删除建立的前驱节点连接
            }
        } else {
            ans = append(ans, cur1.Val) // 到达最左侧
        }
        cur1 = cur1.Right // 利用之前建立的前驱节点返回到根节点
    }
    return ans
}
