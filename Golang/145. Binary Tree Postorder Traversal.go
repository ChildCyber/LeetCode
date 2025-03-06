package leetcode

// 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// 递归
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	// 左 → 右 → 根
	if root != nil {
		ans = append(ans, postorderTraversal(root.Left)...)
		ans = append(ans, postorderTraversal(root.Right)...)
		ans = append(ans, root.Val)
	}
	return ans
}

// 递归-闭包
func postorderTraversalClosure(root *TreeNode) []int {
	var result []int

	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}

	traverse(root)
	return result
}

// 迭代-栈
func postorderTraversalStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 将结果插入到头部
		result = append([]int{node.Val}, result...)

		// 左子节点先入栈（后出栈）
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		// 右子节点后入栈（先出栈）
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}

	return result
}
