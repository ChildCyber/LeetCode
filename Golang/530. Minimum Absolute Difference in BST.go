package leetcode

import "math"

// 二叉搜索树的最小绝对差
// https://leetcode.cn/problems/minimum-absolute-difference-in-bst/

// BFS中序遍历
func getMinimumDifference(root *TreeNode) int {
	// BST的最小差值，必然出现在中序遍历中相邻节点之间
	var prev *TreeNode
	minDiff := math.MaxInt64

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)

		if prev != nil {
			diff := node.Val - prev.Val
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = node

		inorder(node.Right)
	}

	inorder(root)
	return minDiff
}

// 中序遍历到数组
func getMinimumDifferenceWithArray(root *TreeNode) int {
	values := []int{}

	// 中序遍历收集所有值
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		values = append(values, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	// 计算相邻元素的最小差值
	minDiff := math.MaxInt32
	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
