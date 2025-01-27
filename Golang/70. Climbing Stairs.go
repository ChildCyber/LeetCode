package leetcode

// 爬楼梯
// https://leetcode-cn.com/problems/climbing-stairs/
// 动态规划
func climbStairs(n int) int {
	// 状态转移方程 dp[i] = dp[i-2] + dp[i-1]
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 只保留前两个值
func climbStairs1(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
