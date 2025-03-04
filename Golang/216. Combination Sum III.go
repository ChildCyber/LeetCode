package leetcode

// 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/

// 回溯
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var path []int

	// 回溯函数
	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		// 终止条件：选够 k 个数
		if len(path) == k {
			if sum == n {
				// 复制当前组合加入结果
				comb := make([]int, k)
				copy(comb, path)
				res = append(res, comb)
			}
			return
		}

		// 剪枝：如果 sum 已经超了
		if sum > n {
			return
		}

		// 从 start 开始选数
		for i := start; i <= 9; i++ {
			// 如果当前数已经使 sum 超出，没必要继续
			if sum+i > n {
				break
			}
			path = append(path, i)
			dfs(i+1, sum+i)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	dfs(1, 0)
	return res
}
