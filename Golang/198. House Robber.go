package leetcode

// 打家劫舍
// https://leetcode-cn.com/problems/house-robber/
// 动态规划
func rob198(nums []int) int {
	// dp[i]代表抢 nums[0,i] 这个区间内房子的最大值
	// 状态转移方程： dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0], dp[1] = nums[0], max(nums[1], nums[0])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}
	return dp[n-1] // 重点
}

// 动态规划-空间优化
func rob198_1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// 使用两个变量代替整个dp数组
	prev2 := 0 // 相当于 dp[i-2]
	prev1 := 0 // 相当于 dp[i-1]
	for i := 0; i < n; i++ {
		// 计算当前房子的最优解
		current := max(prev1, prev2+nums[i])

		// 更新状态，为下一个房子做准备
		prev2 = prev1
		prev1 = current
	}
	return prev1
}
