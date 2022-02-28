package leetcode

// 最长公共子序列
// https://leetcode-cn.com/problems/longest-common-subsequence
// 动态规划
func longestCommonSubsequence(text1, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
	// 边界条件：
	// 当 i=0 时，text1[0:i] 为空，dp[0][j]=0
	// 当 j=0 时，text2[0:j] 为空，dp[i][0]=0
	dp := make([][]int, m+1) // (m+1)*(n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 { // 两字符相同
				dp[i+1][j+1] = dp[i][j] + 1
			} else { // 两字符不同
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}
