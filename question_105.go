package leetcode

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    if len(preorder) == 1 && len(inorder) == 1 {
        root := &TreeNode{Val: preorder[0]}
        return root
    }

    root := &TreeNode{Val: preorder[0]}
    var left, right *TreeNode

    idx := getPos(inorder, preorder[0])
    leftLen := idx
    rightLen := len(inorder) - idx - 1
    if leftLen > 0 {
        leftPreorder := preorder[1:idx+1]
        leftInorder := inorder[0:idx]
        left = buildTree(leftPreorder, leftInorder)
    }
    if rightLen > 0 {
        rightPreorder := preorder[idx+1:]
        rightInorder := inorder[idx+1:]
        right = buildTree(rightPreorder, rightInorder)
    }
    root.Left = left
    root.Right = right
    return root
}

func getPos(inorder []int, val int) int {
    for idx, v := range inorder {
        if v == val {
            return idx
        }
    }
    return -1
}
