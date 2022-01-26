package leetcode

import "math"

// 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 递归
func isValidBST(root *TreeNode) bool {
	return isValidbst(root, math.MinInt64, math.MaxInt64)
}

func isValidbst(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValidbst(root.Left, lower, root.Val) && isValidbst(root.Right, root.Val, upper)
	//v := float64(root.Val)
	//return v < max && v > min && isValidbst(root.Left, min, v) && isValidbst(root.Right, v, max)
}
