package leetcode

// 目标和
// https://leetcode.cn/problems/target-sum/
// 回溯
func findTargetSumWays(nums []int, target int) (count int) {
	// 使用回溯的方法遍历所有的表达式，回溯过程中维护一个计数器 count，当遇到一种表达式的结果等于目标数 target 时，将 count 的值加 1。
	// 遍历完所有的表达式之后，即可得到结果等于目标数 target 的表达式的数目。
	var backtrack func(int, int)
	backtrack = func(index, sum int) {
		if index == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(index+1, sum+nums[index])
		backtrack(index+1, sum-nums[index])
	}

	backtrack(0, 0)
	return
}
