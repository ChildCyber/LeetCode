package leetcode

// 路径总和 II
// https://leetcode-cn.com/problems/path-sum-ii/
// lc 257和lc 112的组合增强版

// DFS+回溯
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}

	// 递归的三要素：
	// 1. 终止条件
	// 当前节点为null就终止
	// 2. 递归函数做的事情是什么以及如何去做
	// 把当前节点加入到record路径中。
	// 当到达叶子节点时，如果路径和等于目标路径，需要把record加入到结果集合中，然后返回。
	//   返回前需要把当前节点从record路径中移除。
	// 重复利用递归函数。
	// 函数最终返回前需要把当前节点从record路径中移除。
	// 3. 如何重复利用递归函数
	// 递归调用当前函数，从左右子树中去寻找目标路径和。
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, leftNum int) {
		if node == nil {
			return
		}

		leftNum -= node.Val
		path = append(path, node.Val) // 当前节点入队

		// 叶子节点检查
		if node.Left == nil && node.Right == nil {
			if leftNum == 0 { // 当前节点为叶子节点，且路径和等于targetSum
				ans = append(ans, append([]int(nil), path...)) // 该路径为答案之一
			}
		} else {
			// 非叶子节点，继续递归
			dfs(node.Left, leftNum)
			dfs(node.Right, leftNum)
		}

		path = path[:len(path)-1] // 回溯，当前节点出队
	}

	dfs(root, targetSum)
	return
}
