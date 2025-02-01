package leetcode

import "math"

// 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/
// 递归
func isValidBST(root *TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
}

// BFS：使用栈来模拟中序遍历
func isValidBSTBFS(root *TreeNode) bool {
	stack := []*TreeNode{}
	// 记录前一个访问的节点（用于比较）
	var prev *TreeNode = nil
	// 从根节点开始
	curr := root

	for curr != nil || len(stack) > 0 {
		// 遍历到最左节点
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		// 弹出栈顶节点
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 检查当前节点值是否大于前一个节点值
		if prev != nil && curr.Val <= prev.Val {
			return false
		}

		// 更新前一个节点并转向右子树
		prev = curr
		curr = curr.Right
	}

	return true
}
