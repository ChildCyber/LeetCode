package leetcode

// 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/

// 回溯
func combinationSum3(k int, n int) [][]int {
	// 只使用数字1到9 && 每个数字最多使用一次
	var ans [][]int
	var path []int

	// 回溯函数
	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		// 终止条件是选够k个数
		if len(path) == k {
			if sum == n {
				// 复制当前组合加入结果
				comb := make([]int, k)
				copy(comb, path)
				ans = append(ans, comb)
			}
			return
		}

		// 剪枝：如果sum已经超了
		if sum > n {
			return
		}

		// 从start开始选数
		for i := start; i <= 9; i++ {
			// 如果当前数已经使sum超出，没必要继续
			if sum+i > n {
				break
			}
			// 做选择
			path = append(path, i)
			// 递归
			backtrack(i+1, sum+i)
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	// 从1开始是题目的设定
	backtrack(1, 0)
	return ans
}
