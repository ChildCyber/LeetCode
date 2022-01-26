package leetcode

// 二叉树的后序遍历
// https://leetcode-cn.com/problems/binary-tree-postorder-traversal/

// 递归
func postorderTraversalRec(root *TreeNode) []int {
	res := []int{}
	if root != nil {
		tmp := postorderTraversalRec(root.Left)
		for _, t := range tmp {
			res = append(res, t)
		}
		tmp = postorderTraversalRec(root.Right)
		for _, t := range tmp {
			res = append(res, t)
		}
		res = append(res, root.Val)
	}
	return res
}
