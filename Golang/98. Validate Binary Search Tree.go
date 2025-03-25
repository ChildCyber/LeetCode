package leetcode

import "math"

// 验证二叉搜索树
// https://leetcode-cn.com/problems/validate-binary-search-tree/

// 递归
// 时间复杂度：O(n)
// 空间复杂度：O(h)，递归栈深度
func isValidBST(root *TreeNode) bool {
	var isValidBSTHelper func(*TreeNode, int, int) bool
	isValidBSTHelper = func(root *TreeNode, lower, upper int) bool {
		if root == nil {
			return true
		}
		// 检查当前节点值是否在合法范围内
		if root.Val <= lower || root.Val >= upper {
			return false
		}
		// 递归验证左右子树，更新范围
		return isValidBSTHelper(root.Left, lower, root.Val) && isValidBSTHelper(root.Right, root.Val, upper)
	}

	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

// 中序遍历
func isValidBSTInorder(root *TreeNode) bool {
	var prev *TreeNode
	var traverse func(*TreeNode) bool
	traverse = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		// 左子树
		if !traverse(node.Left) {
			return false
		}
		// 当前节点：检查是否严格递增
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		// 右子树
		return traverse(node.Right)
	}
	return traverse(root)
}

// 迭代中序遍历
// 时间复杂度：O(n)
// 空间复杂度：O(h)，栈的空间
func isValidBSTBFS(root *TreeNode) bool {
	stack := []*TreeNode{}
	// 记录前一个访问的节点（用于比较）
	var prev *TreeNode = nil
	// 从根节点开始
	curr := root

	for curr != nil || len(stack) > 0 {
		// 遍历到最左节点，将左子树全部入栈
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
