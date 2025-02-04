package leetcode

import "sort"

// 组合总和
// https://leetcode-cn.com/problems/combination-sum/

// 回溯
func combinationSum(candidates []int, target int) [][]int {
	// 先排序，方便剪枝
	sort.Ints(candidates)

	var res [][]int
	var path []int

	var dfs func(start, sum int)
	dfs = func(start, sum int) {
		// 如果刚好等于 target，保存一份
		if sum == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		// 如果超过 target，直接返回
		if sum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 剪枝：如果当前数 + sum > target，后面更大就不用看了
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			// 这里传 i 而不是 i+1，因为数字可重复使用
			dfs(i, sum+candidates[i])
			path = path[:len(path)-1] // 撤销选择
		}
	}

	dfs(0, 0)
	return res
}

// 回溯-不含剪枝
func combinationSum1(candidates []int, target int) (ans [][]int) {
	path := []int{}

	var dfs func(target, idx int)
	dfs = func(target, idx int) { // 参数：目标和，candidates数组的第 idx 位
		if idx == len(candidates) { // 所有数字都已经全部选完
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 直接跳过，不选择当前数
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 { // 能够选择
			path = append(path, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	dfs(target, 0)
	return
}
