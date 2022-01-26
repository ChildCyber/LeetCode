package leetcode

// 最大子序和
// https://leetcode-cn.com/problems/maximum-subarray/
// 动态规划
func maxSubArray(nums []int) int {
	// 状态转移方程： dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0) , dp[i] = nums[i] (dp[i-1] ≤ 0)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ { // 从1开始
		// dp的两种情况
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i]) // 最大子序和
	}
	return res
}

// 模拟
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}
