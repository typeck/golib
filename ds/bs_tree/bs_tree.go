package bs_tree

import(
	"github.com/typeck/golib/entity"
)

func insert(node *entity.TreeNode, x int) *entity.TreeNode {
	if node == nil {
		return entity.NewTreeNode()
	}
	if x < node.Val {
		node.Left = insert(node.Left, x)
	}else {
		node.Right = insert(node.Right, x)
	}
	return node
}

func find(node *entity.TreeNode, x int) *entity.TreeNode {
	if node == nil {
		return nil
	}
	if x < node.Val {
		return find(node.Left, x)
	}else {
		return find(node.Right, x)
	}
}

func findMin(node *entity.TreeNode) *entity.TreeNode {
	if node == nil || node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

func remove(node *entity.TreeNode, x int) *entity.TreeNode {
	if node == nil {
		return nil
	}
	var tmp *entity.TreeNode
	if x < node.Val {
		node.Left = remove(node.Left, x)
	}else if x > node.Val {
		node.Right = remove(node.Right, x)
		
	 //如果node就是要删除的节点
	}else {
		//如果要删除的节点左右子树都存在，则找到右子树的最小节点与之替换，并删除最小节点
		if node.Left != nil && node.Right != nil {
			tmp = findMin(node.Right)
			node.Val = tmp.Val
			node.Right = remove(node.Right, node.Val)
		 // 处理要删除节点是叶节点，或者只有一棵子树的情况
		}else {
			if node.Left == nil {
				return node.Right
			}
			if node.Right == nil {
				return node.Left
			}
		}
	}
	return node
}
