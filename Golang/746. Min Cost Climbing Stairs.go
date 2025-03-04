package leetcode

// 使用最小花费爬楼梯
// https://leetcode.cn/problems/min-cost-climbing-stairs

// 动态规划
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n)
	dp[0], dp[1] = cost[0], cost[1]
	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-2], dp[i-1])
	}
	return min(dp[n-2], dp[n-1])
}

// 动态规划-空间优化
func minCostClimbingStairsOptimized(cost []int) int {
	n := len(cost)
	if n <= 1 {
		return 0
	}

	// 只用两个变量代替dp数组
	prev2, prev1 := 0, 0 // dp[0] 和 dp[1]

	for i := 2; i <= n; i++ {
		// 计算当前dp[i]
		current := min(prev1+cost[i-1], prev2+cost[i-2])
		// 更新前两个状态
		prev2, prev1 = prev1, current
	}

	return prev1
}
