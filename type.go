package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}
