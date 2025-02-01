package leetcode

// 路径总和 III
// https://leetcode.cn/problems/path-sum-iii/
// DFS 双重递归
// 时间复杂度：O(n²)
func pathSum437(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 以当前节点为起点的路径数量
	countFromRoot := rootSum(root, targetSum)

	// 递归处理左子树和右子树（把它们作为新的起点）
	countFromLeft := pathSum437(root.Left, targetSum)
	countFromRight := pathSum437(root.Right, targetSum)

	return countFromRoot + countFromLeft + countFromRight
}

// 以节点 root 为起点向下且满足路径总和为 targetSum 的路径数目
func rootSum(root *TreeNode, targetSum int) (ans int) {
	if root == nil {
		return
	}

	val := root.Val
	if val == targetSum {
		ans++
	}
	// 对左子树和右子树进行递归搜索，更新剩余目标值
	ans += rootSum(root.Left, targetSum-val)
	ans += rootSum(root.Right, targetSum-val)
	return
}

// 前缀和+哈希表
// 思路类似560题
func pathSum437Prefix(root *TreeNode, targetSum int) int {
	prefixSum := make(map[int]int)
	prefixSum[0] = 1 // 空路径的前缀和为0

	return dfs437(root, 0, targetSum, prefixSum)
}

func dfs437(node *TreeNode, currSum, targetSum int, prefixSum map[int]int) int {
	if node == nil {
		return 0
	}

	// 更新当前前缀和
	currSum += node.Val

	// 计算满足条件的路径数量
	count := prefixSum[currSum-targetSum]

	// 将当前前缀和加入哈希表
	prefixSum[currSum]++

	// 递归遍历左右子树
	count += dfs437(node.Left, currSum, targetSum, prefixSum)
	count += dfs437(node.Right, currSum, targetSum, prefixSum)

	// 回溯：移除当前前缀和
	prefixSum[currSum]--

	return count
}
