package leetcode

import "strconv"

// 二叉树的所有路径
// https://leetcode-cn.com/problems/binary-tree-paths/
// 递归
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}

	res := []string{}
	// 非叶子节点
	// 左子树
	tmpLeft := binaryTreePaths(root.Left)
	for i := 0; i < len(tmpLeft); i++ {
		// 当前节点加上递归返回的左子树
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpLeft[i])
	}
	// 右子树
	tmpRight := binaryTreePaths(root.Right)
	for i := 0; i < len(tmpRight); i++ {
		// 当前节点加上递归返回的右子树
		res = append(res, strconv.Itoa(root.Val)+"->"+tmpRight[i])
	}

	return res
}
