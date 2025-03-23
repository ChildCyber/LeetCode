package leetcode

import "sort"

// 组合总和
// https://leetcode-cn.com/problems/combination-sum/

// 回溯+排序
func combinationSum(candidates []int, target int) [][]int {
	// 先排序，方便剪枝
	sort.Ints(candidates)

	var res [][]int
	var path []int

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		// 找到满足条件的组合
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

			// 做选择
			path = append(path, candidates[i])
			// 递归：这里传 i 而不是 i+1，因为数字可重复使用
			backtrack(i, sum+candidates[i])
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}

// 回溯-不含剪枝
func combinationSum1(candidates []int, target int) (ans [][]int) {
	path := []int{}

	var backtrack func(target, idx int)
	backtrack = func(target, idx int) { // 参数：目标和，candidates数组的第 idx 位
		if idx == len(candidates) { // 所有数字都已经全部选完
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 直接跳过，不选择当前数
		backtrack(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			// 做选择
			path = append(path, candidates[idx])
			// 递归
			backtrack(target-candidates[idx], idx)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	backtrack(target, 0)
	return
}
