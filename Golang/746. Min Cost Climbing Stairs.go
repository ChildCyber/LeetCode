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
