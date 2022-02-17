package leetcode

import "sort"

// 组合总和
// https://leetcode-cn.com/problems/combination-sum/
// 回溯
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	c, res := []int{}, [][]int{}
	sort.Ints(candidates) // 排序
	// target:目标和
	// index:candidates数组的第 idx 位
	// c:当前组合的数的数组
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 { // 先判断目标和
		if target == 0 {
			b := make([]int, len(c))
			copy(b, c)
			*res = append(*res, b)
		}
		return
	}
	for i := index; i < len(nums); i++ {
		if nums[i] > target { // 这里可以剪枝优化
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res) // 注意这里迭代的时候 index 依旧不变,因为一个元素可以取多次
		c = c[:len(c)-1]
	}
}

func combinationSum1(candidates []int, target int) (ans [][]int) {
	comb := []int{}

	var dfs func(target, idx int)
	dfs = func(target, idx int) { // 参数：目标和，candidates数组的第 idx 位
		if idx == len(candidates) { // 所有数字都已经全部选完
			return
		}

		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}

		// 直接跳过，不选择当前数
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 { // 能够选择
			comb = append(comb, candidates[idx]) // 入
			dfs(target-candidates[idx], idx)     // 递归
			// 回溯
			comb = comb[:len(comb)-1] // 出
		}
	}

	dfs(target, 0)
	return
}
