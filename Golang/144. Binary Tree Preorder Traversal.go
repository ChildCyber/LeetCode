package leetcode

// 二叉树前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// 递归
func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	// 根 → 左 → 右
	if root != nil {
		ans = append(ans, root.Val)
		ans = append(ans, preorderTraversal(root.Left)...)
		ans = append(ans, preorderTraversal(root.Right)...)
	}
	return ans
}

// 迭代-栈
func preorderTraversalStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		// 弹出栈顶节点
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 访问当前节点
		ans = append(ans, node.Val)

		// 先右后左入栈（保证出栈时先左后右）
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return ans
}
