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

// 迭代
func inorderTraversalIter(root *TreeNode) (ans []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil { // 到达最左叶子节点
			stack = append(stack, root) // 将当前到达的节点入栈
			root = root.Left
		}
		// 此时root为空，说明上一步的root没有左子树
		// 1. 执行左出栈。因为此时root为空，导致root.right一定为空
		// 2. 执行下一次外层while代码块，根出栈。此时root.right可能存在
		// 3a. 若root.right存在，右入栈，再出栈
		// 3b. 若root.right不存在，重复步骤2
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1] // 出栈，记录当前节点
		ans = append(ans, root.Val)
		root = root.Right
	}
	return
}
