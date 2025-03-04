package leetcode

// 二叉树中的最长交错路径
// https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/

// DFS
func longestZigZag(root *TreeNode) int {
	maxLen := 0

	var dfs func(node *TreeNode, dir int, length int)
	dfs = func(node *TreeNode, dir int, length int) {
		if node == nil {
			return
		}
		if length > maxLen {
			maxLen = length
		}

		if dir == 0 { // 上一步是往左，所以下一步必须往右
			dfs(node.Right, 1, length+1)
			dfs(node.Left, 0, 1) // 重启路径
		} else { // 上一步是往右，下一步必须往左
			dfs(node.Left, 0, length+1)
			dfs(node.Right, 1, 1) // 重启路径
		}
	}

	dfs(root.Left, 0, 1)  // 从 root 往左出发
	dfs(root.Right, 1, 1) // 从 root 往右出发

	return maxLen
}
