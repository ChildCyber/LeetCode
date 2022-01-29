package leetcode

// 二叉树的最大深度
// https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 遍历根节点的左孩子的高度和根节点右孩子的高度，取出两者的最大值再加一即为总高度
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
