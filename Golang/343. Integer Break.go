package leetcode

// 整数拆分，剪绳子
// https://leetcode-cn.com/problems/integer-break/
// 动态规划
func integerBreak(n int) int {
	dp := make([]int, n+1)
	// dp[i] 表示将正整数 i 拆分成至少两个正整数的和之后，这些正整数的最大乘积
	// 0 不是正整数，1 是最小的正整数，0 和 1 都不能拆分，因此 dp[0]=dp[1]=0
	for i := 2; i <= n; i++ {
		// 当 i≥2 时，假设对正整数 i 拆分出的第一个正整数是 j（1≤j<i），则有以下两种方案：
		// 将 i 拆分成 j 和 i−j 的和，且 i−j 不再拆分成多个正整数，此时的乘积是 j×(i−j)；
		// 将 i 拆分成 j 和 i−j 的和，且 i−j 继续拆分成多个正整数，此时的乘积是 j×dp[i−j]。
		for j := 1; j < i; j++ {
			// 状态转移方程：dp[i] = max(dp[i], j * (i - j), j * dp[i-j])
			dp[i] = max(dp[i], j*max(dp[i-j], i-j))
		}
	}
	return dp[n]
}
