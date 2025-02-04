package leetcode

// 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/

// 回溯
func partition131(s string) [][]string {
	ans := make([][]string, 0)
	path := make([]string, 0)

	var isPalindrome func(int, int) bool
	isPalindrome = func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			ans = append(ans, append([]string(nil), path...))
			return
		}

		for end := start; end < len(s); end++ {
			if isPalindrome(start, end) {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return ans
}

// 回溯+动态规划预处理
func partition131DP(s string) (ans [][]string) {
	n := len(s)

	// 预处理回文矩阵
	// dp(i,j) 表示 s[i..j] 是否为回文
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true // 单个字符一定是回文串
		}
	}
	for i := n - 1; i >= 0; i-- { // 剪枝
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1] // 首尾字符相等且中间字符为回文串
		}
	}

	// 回溯枚举所有可能的分割方法并进行判断
	path := []string{}
	var dfs func(int)
	dfs = func(i int) { // 当前已搜索到第i个字符
		if i == n { // 已经到最后一个字符
			ans = append(ans, append([]string(nil), path...))
			return
		}

		for j := i; j < n; j++ {
			if dp[i][j] { // s[i..j]为回文串
				path = append(path, s[i:j+1]) // 选择当前字符
				dfs(j + 1)
				path = path[:len(path)-1] // 回溯，不选
			}
		}
	}

	dfs(0)
	return
}
