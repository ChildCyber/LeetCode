package leetcode

import (
	"strconv"
	"strings"
)

// 二叉树的所有路径
// https://leetcode-cn.com/problems/binary-tree-paths/

// 递归+回溯
func binaryTreePaths(root *TreeNode) []string {
	ans := []string{}
	path := []string{}

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 添加当前节点到路径
		path = append(path, strconv.Itoa(node.Val))

		// 叶子节点，构建路径字符串
		if node.Left == nil && node.Right == nil {
			ans = append(ans, strings.Join(path, "->"))
		} else {
			// 递归遍历左右子树
			dfs(node.Left)
			dfs(node.Right)
		}

		// 回溯
		path = path[:len(path)-1]
	}

	dfs(root)
	return ans
}
