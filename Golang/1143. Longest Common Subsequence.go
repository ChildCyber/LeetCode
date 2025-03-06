package leetcode

// 最长公共子序列
// https://leetcode-cn.com/problems/longest-common-subsequence/

// 暴力
// 时间复杂度：O(2^(m+n))
// 空间复杂度：O(m+n)
func longestCommonSubsequenceBrute(text1, text2 string) int {
	// 思路：穷举所有可能的匹配选择
	// 对于两个字符串的每个位置，只有两种选择：
	//   如果当前字符匹配：必须选择匹配，然后继续处理后面的字符
	//   如果当前字符不匹配：有两种选择：
	//     跳过text1的当前字符，继续匹配
	//     跳过text2的当前字符，继续匹配
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 基准情况：任一字符串遍历完成
		if i == len(text1) || j == len(text2) {
			return 0
		}
		// 当前字符匹配：必须选择匹配
		if text1[i] == text2[j] {
			return 1 + dfs(i+1, j+1)
		}
		// 当前字符不匹配：两种选择
		return max(dfs(i+1, j), dfs(i, j+1))
	}
	return dfs(0, 0)
}

// 动态规划
// 时间复杂度：O(m×n)
// 空间复杂度：O(m×n)
func longestCommonSubsequence(text1, text2 string) int {
	// 暴力递归是通过决策树搜索最优匹配路径，而不是生成所有子串组合。动态规划的本质就是给这种递归加上记忆化，避免重复计算
	m, n := len(text1), len(text2)
	// 状态定义：dp[i][j] 表示 text1[0:i] 和 text2[0:j] 的最长公共子序列的长度
	//   其中i和j分别表示第一个字符串和第二个字符串的位置索引
	// 状态转移方程：根据字符是否匹配，有两种情况
	//   情况1：当前字符匹配（text1[i-1] == text2[j-1]）：dp[i][j] = dp[i-1][j-1] + 1
	//   情况2：当前字符不匹配（text1[i-1] != text2[j-1]）：dp[i][j] = max(dp[i-1][j], dp[i][j-1])
	// 边界条件：任何一个字符串为空时，最长公共子序列长度都为0
	//   当 i=0 时，text1[0:i] 为空，dp[0][j]=0
	//   当 j=0 时，text2[0:j] 为空，dp[i][0]=0
	dp := make([][]int, m+1) // 大小为(m+1)*(n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// 填充DP表
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				// 字符匹配，长度+1
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 字符不匹配，取上方或左方的最大值
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}

// 动态规划-空间优化
func longestCommonSubsequenceOptimized(text1 string, text2 string) int {
	m, n := len(text1), len(text2)

	// 使用一维数组进行空间优化
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0 // 保存dp[i-1][j-1]的值
		for j := 1; j <= n; j++ {
			temp := dp[j] // 保存当前值，作为下一轮的prev
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = temp
		}
	}

	return dp[n]
}
