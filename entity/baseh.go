package entity

// 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode() *TreeNode {
	return &TreeNode{}
}