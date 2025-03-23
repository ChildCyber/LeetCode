package leetcode

import "sort"

// 全排列 II
// https://leetcode-cn.com/problems/permutations-ii/

// 回溯+排序
func permuteUnique(nums []int) [][]int { // 可包含重复数字的序列
	if len(nums) == 0 {
		return [][]int{}
	}

	ans := [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑

	// path记录已经选了哪些数
	// used记录下标为i的nums[i]是否已经使用
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		// 找到完整排列
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		// 每次都从头开始选择（但要排除已使用的）
		for i := 0; i < len(nums); i++ {
			// 跳过已使用的元素
			if used[i] {
				continue
			}

			// 关键去重逻辑：同一层遇到相同元素跳过
			// i > 0 保证不越界
			// nums[i] == nums[i-1] 当前元素与前一个相同
			// !used[i-1] 前一个相同元素未被使用（说明在同一层）
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 做选择
			used[i] = true
			path = append(path, nums[i])

			// 递归
			backtrack(path, used)

			// 撤销选择（回溯）
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]int{}, make([]bool, len(nums)))

	return ans
}
