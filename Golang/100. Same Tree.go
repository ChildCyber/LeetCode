package leetcode

// 相同的树
// https://leetcode-cn.com/problems/same-tree/
// 递归
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil { // 都为空
		return true
	} else if p != nil && q != nil { // 都非空
		if p.Val != q.Val { // 值不同
			return false
		}
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right) // 左右子树是否相同
	} else {
		return false
	}
}
