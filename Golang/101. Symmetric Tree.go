package leetcode

// 对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/

// 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil { // 空树
		return true
	}
	// 判断左右子树是否对称
	return isSameTree(invertTree(root.Left) /*翻转左子树*/, root.Right) // lc 226, 100
}

// 左子树的左子树与右子树的右子树对称，且左子树的右子树与右子树的左子树对称。
func isSymmetricMirror(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// 辅助函数：判断两棵树是否镜像对称
func isMirror(left, right *TreeNode) bool {
	// 如果两个节点都为空，则对称
	if left == nil && right == nil {
		return true
	}
	// 如果只有一个节点为空，则不对称
	if left == nil || right == nil {
		return false
	}
	// 如果节点值不相等，则不对称
	if left.Val != right.Val {
		return false
	}

	// 递归检查左子树的左节点和右子树的右节点是否对称，以及左子树的右节点和右子树的左节点是否对称
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}
