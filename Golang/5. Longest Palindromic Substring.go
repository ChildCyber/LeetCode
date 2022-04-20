package leetcode

// 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 暴力
// 穷举出所有子字符串；从字符串首字符比较到末尾，记录最长回文长度
func longestPalindromeForce(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	begin, maxLen := 0, 1
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ { // i~j为子字符串
			if j-i+1 > maxLen && isPalindrome5(s, i, j) { // i~j为回文串且长度超过当前记录最大长度
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

func isPalindrome5(s string, i, j int) bool {
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 动态规划
// 时间复杂度 O(n^2),空间复杂度 O(n^2)
func longestPalindrome(s string) string { // 返回子串
	ans, dp := "", make([][]bool, len(s))
	// 定义 dp[i][j] 表示从字符串第 i 个字符到第 j 个字符这一段子串是否为回文串
	// 由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串
	// 只有 s[i+1:j−1] 是回文串，并且 s 的第 i 和 j 个字母相同时，s[i:j] 才会是回文串
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s)) // dp[i][j] 表示 s[i..j] 是否是回文串
	}

	// 状态转移方程：dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1]) ,
	// j - i < 3 的情况：
	// j - i == 1 的时候，即只有 2 个字符的情况，只需要判断这 2 个字符是否相同即可
	// j - i == 2 的时候，即只有 3 个字符的情况，只需要判断除去中心以外对称的 2 个字符是否相等

	// 每次循环动态维护保存最⻓回文串
	// i为子串起点，j为子串终点
	for i := len(s) - 1; i >= 0; i-- { // i从后向前
		for j := i; j < len(s); j++ { // j从前向后，j初始值为i
			// s 的第 i 和 j 个字母相同 且 （一个或两个字符串 或 s[i+1..j-1] 为回文串）
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && (ans == "" || j-i+1 > len(ans)) { // 子串为空或当前子串长度大于最大子串长度，更新当前最大子串长度
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

// 中心扩散法
// 时间复杂度 O(n^2),空间复杂度 O(1)

// Manacher算法
// 时间复杂度 O(n),空间复杂度 O(n)
