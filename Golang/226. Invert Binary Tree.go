package leetcode

// 翻转二叉树
// https://leetcode-cn.com/problems/invert-binary-tree/
// 递归
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)                         // 翻转以左子树为根节点的树
	invertTree(root.Right)                        // 翻转以右子树为根节点的树
	root.Left, root.Right = root.Right, root.Left // 翻转左右子树
	return root
}
