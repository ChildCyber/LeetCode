package leetcode

// 路径总和
// https://leetcode-cn.com/problems/path-sum/

// 递归
// 时间复杂度：O(n)，每个节点最多访问一次
// 空间复杂度：O(h)，h为树的高度
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
