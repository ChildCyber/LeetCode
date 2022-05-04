package leetcode

import (
	"sort"
)

// 子集 II
// https://leetcode.cn/problems/subsets-ii/
// 回溯
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)

	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) { // 已经选到nums的最后一个元素
			ans = append(ans, append([]int(nil), t...))
			return
		}

		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] { // 未选择上一个元素，且上个元素与当前元素相同，跳过
			return
		}
		t = append(t, nums[cur]) // 和上个元素不同，选择当前元素
		dfs(true, cur+1)
		t = t[:len(t)-1] // 回溯，不选择当前元素
	}

	dfs(false, 0)
	return
}

func subsetsWithDup1(nums []int) [][]int {
	c, ans := []int{}, [][]int{}
	sort.Ints(nums) // 去重的关键

	var generateSubsetsWithDup func(int, int)
	generateSubsetsWithDup = func(k int, start int) {
		if len(c) == k {
			ans = append(ans, append([]int(nil), c...))
			return
		}

		for i := start; i < len(nums)-(k-len(c))+1; i++ { // i的大小不会超过nums长度
			if i > start && nums[i] == nums[i-1] { // 这里是去重的关键逻辑，本次不取重复数字，下次循环可能会取重复数字
				continue
			}
			c = append(c, nums[i])
			generateSubsetsWithDup(k, i+1)
			c = c[:len(c)-1]
		}
		return
	}

	for k := 0; k <= len(nums); k++ { // 子集的长度为k
		generateSubsetsWithDup(k, 0)
	}
	return ans
}
