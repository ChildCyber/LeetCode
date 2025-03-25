package leetcode

import "math"

// 二叉树中的最大路径和
// https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/

// 递归
func maxPathSum(root *TreeNode) int {
	// 对于树中的任意一个节点，经过该节点的最大路径和只有三种可能情况：
	// 1. 只有当前节点本身
	// 2. 当前节点 + 左子树的最大路径
	// 3. 当前节点 + 右子树的最大路径
	// 当考虑全局最大路径和时，还有第四种情况：
	// 4. 左子树最大路径 + 当前节点 + 右子树最大路径（拱形路径）
	if root == nil {
		return 0
	}

	// 初始化最大值为最小整数
	maxSum := math.MinInt32
	// 定义递归函数
	// 返回值：从该节点向下延伸的最大路径和（只能选择左或右一支）
	// 副作用：在递归过程中更新全局最大路径和
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int { // 该函数计算经过该节点的最大路径和
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值，只有在最大贡献值大于0时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 计算经过当前节点的最大路径和（拱形路径）
		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新全局最大值
		maxSum = max(maxSum, priceNewPath)

		// 返回当前节点的最大贡献值（只能选择一边）
		return node.Val + max(leftGain, rightGain)
	}

	// 从根节点开始递归
	maxGain(root)

	return maxSum
}
