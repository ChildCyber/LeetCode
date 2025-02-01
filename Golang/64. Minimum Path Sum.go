package leetcode

// 最小路径和
// https://leetcode-cn.com/problems/minimum-path-sum/
// 动态规划
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	// 状态定义：dp[i][j] 表示从起点 (0, 0) 走到位置 (i, j) 的最小路径和
	// 状态转移方程：	dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
	// 解释： 到达 (i, j) 的最小路径和，等于从上面来的路径和与从左边来的路径和中的较小值，再加上当前格子的值
	// 初始化：
	// 第一行 dp[0][j] = dp[0][j-1] + grid[0][j]：因为只能从左边过来
	// 第一列 dp[i][0] = dp[i-1][0] + grid[i][0]：因为只能从上面过来
	// 起点 dp[0][0] = grid[0][0]
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, columns)
	}

	// 初始化起点
	dp[0][0] = grid[0][0]
	// 初始化第一列
	for i := 1; i < rows; i++ { // 第0列，除去dp[0][0]
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	// 初始化第一行
	for j := 1; j < columns; j++ { // 第0行，除去dp[0][0]
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// 动态规划填表
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

// 动态规划-一维数组
func minPathSum1(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]

	// 初始化第一行
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		// 每行第一个元素
		dp[0] += grid[i][0]
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j]
		}
	}
	return dp[n-1]
}

// 原地DP，无辅助空间
func minPathSumOptimized(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
