package leetcode

import (
	"sort"
)

// 子集 II
// https://leetcode.cn/problems/subsets-ii/

// 子集的定义： 原集合中任意元素组成的集合，包括空集。

// 回溯+排序
// 回溯-显式
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums) // 关键：排序，便于同层去重

	var ans [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(nums) { // 已经选到nums中最后一个元素
			ans = append(ans, append([]int(nil), path...))
			return
		}

		// 决策一：选当前元素（和重复无关）
		path = append(path, nums[start])
		// 递归
		backtrack(start + 1)
		// 撤销选择（回溯）
		path = path[:len(path)-1]

		// 决策二：不选当前元素（必须去重）
		for start+1 < len(nums) && nums[start+1] == nums[start] {
			start++
		}
		backtrack(start + 1)
	}

	backtrack(0)
	return ans
}

// 回溯-隐式（不选当前元素的决策隐藏在for循环和递归的结构里，没进循环就是不选，递归返回时天然保留了“不选”的路径）
func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums) // 关键：排序，便于同层去重

	var ans [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		// 每到一层就把当前path记录为一个子集（注意要拷贝）
		tmp := make([]int, len(path))
		copy(tmp, path)
		ans = append(ans, tmp)

		for i := start; i < len(nums); i++ {
			// 核心去重：如果和前一个相同且在同一层（i > start），跳过
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			// 选择nums[i]
			path = append(path, nums[i])
			backtrack(i + 1)
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return ans
}
