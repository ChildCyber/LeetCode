package leetcode

// 删除二叉搜索树中的节点
// https://leetcode.cn/problems/delete-node-in-a-bst/

// 递归
func deleteNode450(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	// 找到要删除的节点
	if key < root.Val {
		root.Left = deleteNode450(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode450(root.Right, key)
	} else {
		// 情况1: 要删除的节点是叶子节点或只有一个子节点
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// 情况2: 要删除的节点有两个子节点
		// 找到右子树中的最小节点
		minNode := findSubTreeMin(root.Right)
		// 用最小节点的值替换当前节点的值
		root.Val = minNode.Val
		// 删除右子树中的最小节点
		root.Right = deleteNode450(root.Right, minNode.Val)
	}

	return root
}

// 辅助函数：找到子树中的最小节点
func findSubTreeMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
