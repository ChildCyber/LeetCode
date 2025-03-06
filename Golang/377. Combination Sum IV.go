package leetcode

// 组合总和 Ⅳ
// https://leetcode.cn/problems/combination-sum-iv/

// 暴力
func combinationSum4Brute(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	c, ans := []int{}, 0
	findCombinationSum4(nums, target, 0, c, &ans)
	return ans
}

func findCombinationSum4(nums []int, target, index int, c []int, ans *int) {
	if target <= 0 {
		if target == 0 {
			*ans++
		}
		return
	}
	for i := 0; i < len(nums); i++ {
		c = append(c, nums[i])
		findCombinationSum4(nums, target-nums[i], i, c, ans)
		c = c[:len(c)-1]
	}
}

// 动态规划
func combinationSum4(nums []int, target int) int {
	// 用 dp[x] 表示选取的元素之和等于 x 的方案数，目标是求 dp[target]
	dp := make([]int, target+1)
	dp[0] = 1                      // 初始化 dp[0]=1
	for i := 1; i <= target; i++ { // 遍历 i 从 1 到 target
		for _, num := range nums { // 对于每个 i，遍历数组 nums 中的每个元素 num，当 num≤i 时，将 dp[i−num] 的值加到 dp[i]
			if num <= i {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
