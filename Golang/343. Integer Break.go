package leetcode

// 整数拆分，剪绳子
// https://leetcode-cn.com/problems/integer-break/
func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			// 状态转移方程：dp[i] = max(dp[i], j * (i - j), j * dp[i-j])
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}
