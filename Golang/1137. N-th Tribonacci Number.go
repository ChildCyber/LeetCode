package leetcode

// 第 N 个泰波那契数
// https://leetcode.cn/problems/n-th-tribonacci-number/

// 动态规划
func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1], dp[2] = 0, 1, 1
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}

// 动态规划-空间优化
func tribonacciOptimized(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	// 初始化前三个值
	a, b, c := 0, 1, 1

	// 从第3项开始计算到第n项
	for i := 3; i <= n; i++ {
		// 计算下一项
		next := a + b + c
		// 更新前三个值
		a, b, c = b, c, next
	}

	return c
}
