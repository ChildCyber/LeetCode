package leetcode

import "sort"

// 全排列 II
// https://leetcode-cn.com/problems/permutations-ii/
// 回溯
func permuteUnique(nums []int) [][]int { // 可包含重复数字的序列
	if len(nums) == 0 {
		return [][]int{}
	}

	// 状态变量
	// path:已经选了哪些数，数组
	// depth:递归到第几层
	// used布尔数组，下标为i的nums[i]是否已经使用
	// res:结果
	used, path, res := make([]bool, len(nums)), []int{}, [][]int{}
	sort.Ints(nums) // 这里是去重的关键逻辑
	generatePermutation47(nums, 0, path, &res, &used)
	return res
}

func generatePermutation47(nums []int, depth int, path []int, res *[][]int, used *[]bool) {
	if depth == len(nums) {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !(*used)[i] {
			if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] { // 这里是去重的关键逻辑
				continue
			}
			(*used)[i] = true
			path = append(path, nums[i])
			generatePermutation47(nums, depth+1, path, res, used)
			// 回溯
			path = path[:len(path)-1]
			(*used)[i] = false
		}
	}
	return
}
