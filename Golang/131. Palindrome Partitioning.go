package leetcode

// 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/
// 回溯 + 动态规划预处理
func partition131(s string) (ans [][]string) {
	n := len(s)

	//  dp(i,j) 表示 s[i..j] 是否为回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true // 单个字符是回文串
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}

	splits := []string{}
	// 回溯枚举所有可能的分割方法并进行判断
	var dfs func(int)
	dfs = func(i int) { // 当前搜索到字符串的第 i 个字符
		if i == n { // 已经到最后一个字符串
			ans = append(ans, append([]string(nil), splits...))
			return
		}

		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1]) // 选择当前字符
				dfs(j + 1)
				splits = splits[:len(splits)-1] // 不选
			}
		}
	}

	dfs(0)
	return
}
