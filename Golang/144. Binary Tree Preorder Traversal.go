package leetcode

// 二叉树前序遍历
// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// 递归
func preorderTraversalRec1(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		res = append(res, root.Val)
		tmp := preorderTraversalRec1(root.Left)
		for _, t := range tmp {
			res = append(res, t)
		}
		tmp = preorderTraversalRec1(root.Right)
		for _, t := range tmp {
			res = append(res, t)
		}
	}
	return res
}

// 栈
func preorderTraversalStack(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, res := []*TreeNode{}, []int{}
	stack = append(stack, root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node != nil {
			res = append(res, node.Val)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}
