package leetcode

// 不同路径
// https://leetcode-cn.com/problems/unique-paths/
// 暴力递归
func uniquePathBrute(m int, n int) int {
	var recursive func(int, int, int, int) int
	recursive = func(m int, n int, i int, j int) int { // 求解从 (0,0) 到右下角的路径数量
		if i == m-1 || j == n-1 {                      // 当已经处于矩阵的最后一行或者最后一列，即只一条路可以走
			return 1
		}
		return recursive(m, n, i+1, j) + recursive(m, n, i, j+1)
	}
	return recursive(m, n, 0, 0)
}

// 动态规划
func uniquePaths(m int, n int) int {
	// 状态定义：dp[i][j] 表示从起点 (0, 0) 走到位置 (i, j) 的不同路径数
	// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
	// 到达 (i, j) 的路径只能从左边 (i, j-1) 或上面 (i-1, j) 过来，所以路径数就是这两个位置的路径数之和
	dp := make([][]int, m) // 行
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 动态规划-空间优化
func uniquePathsOptimized(m int, n int) int {
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
