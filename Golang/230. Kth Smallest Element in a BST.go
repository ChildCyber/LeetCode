package leetcode

// 二叉搜索树中第K小的元素
// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst
// 中序遍历
func kthSmallest(root *TreeNode, k int) int {
	res, count := 0, 0 // 第K小的元素的值，第几次，count==k时就是答案
	inorder230(root, k, &count, &res)
	return res
}

func inorder230(node *TreeNode, k int, count *int, ans *int) {
	if node != nil {
		inorder230(node.Left, k, count, ans) // 左
		*count++                             // 中
		if *count == k {
			*ans = node.Val
			return
		}
		inorder230(node.Right, k, count, ans) // 右
	}
}
