package leetcode

// 二叉树的最小深度
// https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// 一边为空，只能走另一边
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	// 两边都不空
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 简化版DFS
func minDepthSimple(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepthSimple(root.Left)
	right := minDepthSimple(root.Right)

	// 关键：如果某子树为空，不能直接取min（因为不是叶子节点）
	if root.Left == nil || root.Right == nil {
		return left + right + 1
	}

	return min(left, right) + 1
}

// BFS
func minDepthBFS(root *TreeNode) int {
	// 第一个遇到叶子的层数就是最小深度
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	depth := 1

	for len(queue) > 0 {
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// 找到第一个叶子节点，立即返回
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		depth++
	}

	return depth
}
