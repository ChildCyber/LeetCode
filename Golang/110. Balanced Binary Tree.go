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

// 自底向上的递归
func isBalanced1(root *TreeNode) bool {
	return checkHeight(root) != -1
}

// 在计算高度的同时判断平衡性
func checkHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	// 检查左子树
	leftHeight := checkHeight(node.Left)
	if leftHeight == -1 {
		return -1 // 左子树不平衡
	}

	// 检查右子树
	rightHeight := checkHeight(node.Right)
	if rightHeight == -1 {
		return -1 // 右子树不平衡
	}

	// 检查当前节点是否平衡
	if abs(leftHeight-rightHeight) > 1 {
		return -1 // 当前节点不平衡
	}

	// 返回当前节点的高度
	return max(leftHeight, rightHeight) + 1
}
