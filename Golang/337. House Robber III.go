package leetcode

// 打家劫舍 III
// https://leetcode.cn/problems/house-robber-iii/
// 动态规划
func rob337(root *TreeNode) int {
	a, b := dfsTreeRob(root)
	return max(a, b)
}

func dfsTreeRob(root *TreeNode) (a, b int) { // a 代表不打劫，b 代表打劫，所能获取到的最大值
	if root == nil {
		return 0, 0
	}

	l0, l1 := dfsTreeRob(root.Left)
	r0, r1 := dfsTreeRob(root.Right)
	// 当前节点没有被打劫
	// 当 root 不被选中时，root 的左右孩子可以被选中，也可以不被选中。对于 root 的某个具体的孩子 x，它对 root 的贡献是 x 被选中和不被选中情况下权值和的较大值
	tmp0 := max(l0, l1) + max(r0, r1)
	// 当前节点被打劫
	// 当 root 被选中时，root 的左右孩子都不能被选中，故 root 被选中情况下子树上被选中点的最大权值和为 l 和 r 不被选中的最大权值和相加
	tmp1 := root.Val + l0 + r0

	return tmp0, tmp1
}
