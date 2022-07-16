package leetcode

// 二叉树剪枝
// https://leetcode.cn/problems/binary-tree-pruning/
// 递归
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 { // 到达叶子节点，且叶子节点的值为0
		return nil
	}

	return root
}
