package leetcode

// 打家劫舍
// https://leetcode-cn.com/problems/house-robber/
// 动态规划
func rob198(nums []int) int {
	// dp[i]代表抢 nums[0,i] 这个区间内房子的最大值,状态转移方程是 dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// dp[i] 代表抢 nums[0...i] 房子的最大价值
	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1] // 重点
}
