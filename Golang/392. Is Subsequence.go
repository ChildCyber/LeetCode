package leetcode

// 判断子序列
// https://leetcode.cn/problems/is-subsequence/

// 暴力
func isSubsequenceBrute(s string, t string) bool {
	var isSubsequence func(string, string, int, int) bool
	isSubsequence = func(s string, t string, i int, j int) bool {
		if i == len(s) { // s的所有字符都匹配完了
			return true
		}
		if j == len(t) { // t已经遍历完但s还没匹配完
			return false
		}

		// 如果当前字符匹配，两种选择：匹配这个字符或不匹配
		if s[i] == t[j] {
			return isSubsequence(s, t, i+1, j+1)
		}
		return isSubsequence(s, t, i, j+1)
	}

	return isSubsequence(s, t, 0, 0)
}

// 动态规划
func isSubsequenceDP(s string, t string) bool {
	// 定义状态：
	// dp[i][j]表示s的前i个字符是否是t的前j个字符的子序列
	// 状态转移方程：
	// 如果s[i-1] == t[j-1]：dp[i][j] = dp[i-1][j-1]
	// 否则：dp[i][j] = dp[i][j-1]
	// 边界条件：
	// dp[0][j] = True（空字符串是任何字符串的子序列）
	// dp[i][0] = False（非空字符串不可能是空字符串的子序列）
	if len(s) == 0 {
		return true
	}
	if len(t) == 0 {
		return false
	}

	m, n := len(s), len(t)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// 初始化：空字符串是任何字符串的子序列
	for j := 0; j <= n; j++ {
		dp[0][j] = true
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[m][n]
}

// 双指针
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}

	return i == len(s)
}
