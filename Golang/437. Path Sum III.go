package leetcode

// 路径总和 III
// https://leetcode.cn/problems/path-sum-iii/
// DFS
// 以节点 root 为起点向下且满足路径总和为 targetSum 的路径数目
func rootSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}

	val := root.Val
	if val == targetSum {
		ans++
	}
	// 对左子树和右子树进行递归搜索
	ans += rootSum(root.Left, targetSum-val)
	ans += rootSum(root.Right, targetSum-val)
	return
}

func pathSum437(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 以当前节点 root 为目标路径的起点递归向下进行搜索
	ans := rootSum(root, targetSum)
	// 对左子树和右子树进行递归搜索
	ans += pathSum437(root.Left, targetSum)
	ans += pathSum437(root.Right, targetSum)
	return ans
}
