package leetcode

import "sort"

// 组合总和 II
// https://leetcode-cn.com/problems/combination-sum-ii/

// 回溯+排序
// 时间复杂度：O(2^n × n)
// 空间复杂度：O(n) 递归栈深度
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	sort.Ints(candidates) // 去重

	ans := [][]int{}
	var backtrack func(start int, path []int, sum int)
	backtrack = func(start int, path []int, sum int) {
		// 找到满足条件的组合
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return // 只是结束当前递归，不是提前终止搜索
		}

		// 超过目标值，剪枝
		if sum > target {
			return
		}

		// i:=start：不可回头，每次都从上一次选择的后面开始
		for i := start; i < len(candidates); i++ {
			// 关键去重逻辑：同一层遇到相同元素跳过
			// i > start 保证不是第一个元素
			// candidates[i] == candidates[i-1] 当前元素与前一个相同
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// 剪枝：如果当前元素已经使和超过target，后面的更大元素也会超过
			if sum+candidates[i] > target {
				break
			}

			// 做选择
			path = append(path, candidates[i])

			// 递归：从i+1开始，每个数字只能用一次
			backtrack(i+1, path, sum+candidates[i])

			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{}, 0)

	return ans
}
