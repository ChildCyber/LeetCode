package leetcode

// hard
// 编辑距离
// https://leetcode.cn/problems/edit-distance/
// 动态规划
func minDistance(word1, word2 string) int {
	n := len(word1)
	m := len(word2)

	// 有一个字符串为空串
	if n*m == 0 {
		return n + m
	}

	// DP数组
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	// 边界状态初始化
	for i := 0; i < n+1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < m+1; j++ {
		dp[0][j] = j
	}
	// 计算所有DP值
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			left := dp[i-1][j] + 1
			down := dp[i][j-1] + 1
			leftDown := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				leftDown++
			}
			dp[i][j] = min(leftDown, min(left, down))
		}
	}
	return dp[n][m]
}
