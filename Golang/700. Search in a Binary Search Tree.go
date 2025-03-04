package leetcode

// 二叉搜索树中的搜索
// https://leetcode.cn/problems/search-in-a-binary-search-tree/

// 递归
func searchBST(root *TreeNode, val int) *TreeNode {
	// 基准情况：节点为空或找到目标值
	if root == nil || root.Val == val {
		return root
	}

	// 利用BST性质决定搜索方向
	if val < root.Val {
		// 目标值小于当前节点值，在左子树中搜索
		return searchBST(root.Left, val)
	} else {
		// 目标值大于当前节点值，在右子树中搜索
		return searchBST(root.Right, val)
	}
}

// 迭代
func searchBSTIter(root *TreeNode, val int) *TreeNode {
	current := root

	for current != nil {
		if current.Val == val {
			return current
		} else if val < current.Val {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}
