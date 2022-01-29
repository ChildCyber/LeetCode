package leetcode

// 对称二叉树
// https://leetcode-cn.com/problems/symmetric-tree/
// 递归
func isSymetric(root *TreeNode) bool {
	if root == nil { // 空树
		return true
	}
	// 判断左右子树是否对称
	return isSameTree(invertTree(root.Left) /*翻转左子树*/, root.Right) // leetcode 226, 100
}
