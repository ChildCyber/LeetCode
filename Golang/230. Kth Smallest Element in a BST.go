package leetcode

// 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	// 第k小就是中序遍历时访问到的第k个节点
	ans := 0
	count := 0

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 只要遍历到第k个即可停下
		if node == nil || count >= k {
			return
		}

		inorder(node.Left)

		count++
		if count == k {
			ans = node.Val
			return
		}

		inorder(node.Right)
	}

	inorder(root)
	return ans
}
