package leetcode

// 二叉树的中序遍历
// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// dfs
func inorderTraversal(root *TreeNode) []int {
	var result []int
	inorder(root, &result)
	return result
}

func inorder(root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}
