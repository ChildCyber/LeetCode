package leetcode

// 路径总和 II
// https://leetcode-cn.com/problems/path-sum-ii/
// 第 257 题和第 112 题的组合增强版

// DFS
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
func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	path := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, leftNum int) {
		if node == nil {
			return
		}

		leftNum -= node.Val
		path = append(path, node.Val) // 入队
		//defer func() { path = path[:len(path)-1] }() // 出队
		if node.Left == nil && node.Right == nil && leftNum == 0 { // 到达叶子节点，且路径和等于目标路径
			ans = append(ans, append([]int(nil), path...)) // 该路径为答案之一
			return
		}

		dfs(node.Left, leftNum)
		dfs(node.Right, leftNum)
		path = path[:len(path)-1] // 回溯，出队
	}
	dfs(root, targetSum)
	return
}
