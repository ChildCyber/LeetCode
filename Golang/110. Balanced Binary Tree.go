package leetcode

// 平衡二叉树
// https://leetcode-cn.com/problems/balanced-binary-tree/

// 自顶向下的递归
func isBalanced(root *TreeNode) bool { // 判断是否为平衡二叉树
	if root == nil {
		return true
	}

	leftHeight := treeDepth(root.Left)   // 左子树的高度
	rightHeight := treeDepth(root.Right) // 右子树的高度
	// 左右子树的高度相差不超过1且左右子树平衡
	return abs(leftHeight-rightHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

// 获取树的深度
func treeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(treeDepth(root.Left), treeDepth(root.Right)) + 1
}
