package leetcode

import (
	"sort"
)

// 全排列
// https://leetcode-cn.com/problems/permutations/

// 暴力
func permuteBrute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	// idx 用来保存当前位置选择了哪一个下标
	idx := make([]int, n)
	used := make([]bool, n)

	// 一个递归，把多层循环“自动”展开
	var dfs func(pos int)
	dfs = func(pos int) {
		if pos == n {
			// 收集当前排列
			perm := make([]int, n)
			for i := 0; i < n; i++ {
				perm[i] = nums[idx[i]]
			}
			res = append(res, perm)
			return
		}
		// 相当于在第 pos 层循环遍历
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			used[i] = true
			idx[pos] = i
			dfs(pos + 1)
			used[i] = false
		}
	}

	dfs(0)
	return res
}

// 回溯， DFS，树形结构
// 时间复杂度：O(n×n!)，其中 n 为序列的长度
// 空间复杂度：O(n)，其中 n 为序列的长度
func permute(nums []int) [][]int { // 不含重复数字的数组
	// used：布尔数组，记录哪些数字已经被使用
	used, ans := make([]bool, len(nums)), [][]int{}

	// path：数组，记录已经选了哪些数
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		// 如果当前路径长度等于原数组长度，说明找到了一个完整排列
		if len(path) == len(nums) {
			// 创建路径的副本并添加到结果中
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		// 遍历所有数字，依次挑
		for i := 0; i < len(nums); i++ {
			// 如果当前数字已经被使用，跳过
			if used[i] {
				continue
			}

			// 选择当前数字
			used[i] = true
			path = append(path, nums[i])

			// 递归处理下一个位置
			backtrack(path, used)

			// 撤销选择（回溯）
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack([]int{}, used)

	return ans
}

// 回溯-非递归
func permuteIter(nums []int) [][]int {
	sort.Ints(nums) // 排序成最小字典序
	var res [][]int

	temp := make([]int, len(nums))
	copy(temp, nums)
	res = append(res, temp)

	for nextPermutation46(nums) {
		temp = make([]int, len(nums))
		copy(temp, nums)
		res = append(res, temp)
	}
	return res
}

// 下一个排列算法（就地修改 nums）
func nextPermutation46(nums []int) bool {
	// 1. 从后往前找到第一个 nums[i] < nums[i+1]
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return false // 已经是最后一个排列
	}

	// 2. 从后往前找到第一个大于 nums[i] 的数
	j := len(nums) - 1
	for nums[j] <= nums[i] {
		j--
	}

	// 3. 交换 i 和 j
	nums[i], nums[j] = nums[j], nums[i]

	// 4. 反转 i+1 到末尾
	reverse(nums[i+1:])
	return true
}
