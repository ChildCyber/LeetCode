package leetcode

// 最长回文子串
// https://leetcode-cn.com/problems/longest-palindromic-substring/
// 暴力
// 穷举出所有子字符串；从字符串首字符比较到末尾，记录最长回文长度
func longestPalindromeBrute(s string) string {
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

// 中心扩散法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func longestPalindromeCenter(s string) string {
	if len(s) == 0 {
		return ""
	}

	start, end := 0, 0 // 记录最长回文子串的起始和结束索引

	for i := 0; i < len(s); i++ {
		// 以当前字符为中心，扩展查找奇数长度回文
		len1 := expandAroundCenter(s, i, i)
		// 以当前字符和下一个字符之间的空隙为中心，扩展查找偶数长度回文
		len2 := expandAroundCenter(s, i, i+1)

		// 取两种情况中的较大长度
		maxLen := max(len1, len2)

		// 如果找到更长的回文子串，更新起始和结束位置
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

// 从中心向两边扩展，返回回文串的长度
func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	// 返回回文串的长度
	return right - left - 1
}

// 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func longestPalindrome(s string) string { // 返回子串
	// 状态定义：dp[i][j] 表示 s[i:j+1] 是否是回文
	// 由回文串的性质可以得知，回文串去掉一头一尾相同的字符以后，剩下的还是回文串
	// 只有 s[i+1:j−1] 是回文串，并且 s 的第 i 和 j 个字母相同时，s[i:j] 才会是回文串
	// 状态转移方程：dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
	// j - i < 3 的情况：
	// j - i == 1 的时候，即只有 2 个字符的情况，只需要判断这 2 个字符是否相同即可
	// j - i == 2 的时候，即只有 3 个字符的情况，只需要判断除去中心以外对称的 2 个字符是否相等
	ans, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	// 每次循环动态维护保存最⻓回文串
	// i为子串起点，j为子串终点
	for i := len(s) - 1; i >= 0; i-- { // i从后向前
		for j := i; j < len(s); j++ { // j从前向后，j初始值为i
			// s 的第 i 和 j 个字母相同 且 （一个或两个字符串 或 s[i+1..j-1] 为回文串）
			dp[i][j] = s[i] == s[j] && (j-i < 3 || dp[i+1][j-1])
			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && (ans == "" || j-i+1 > len(ans)) { // 子串为空或当前子串长度大于最大子串长度，更新当前最大子串长度
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

func longestPalindromeDP(s string) string { // 返回子串
	length := len(s)
	if length < 2 {
		return s
	}

	ans, dp := "", make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length) // dp[i][j] 表示 s[i..j] 是否是回文串
	}

	// 递推开始
	// 先枚举子串长度
	for l := 2; l <= length; l++ {
		// 枚举左边界，左边界的上限设置可以宽松一些
		for i := 0; i < length; i++ {
			j := l + i - 1
			// 如果右边界越界，就可以退出当前循环
			if j >= length {
				break
			}

			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (ans == "" || j-i+1 > len(ans)) { // 子串为空或当前子串长度大于最大子串长度，更新当前最大子串长度
				ans = s[i : j+1]
			}
		}
	}
	return ans
}

// Manacher算法
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func longestPalindromeManacher(s string) string {
	if len(s) <= 1 {
		return s
	}

	// 预处理字符串，插入特殊字符
	t := preProcess(s)
	n := len(t)
	p := make([]int, n) // 记录每个位置的回文半径

	center, right := 0, 0
	maxCenter, maxLen := 0, 0

	for i := 1; i < n-1; i++ {
		// 利用对称性快速计算初始回文半径
		if i < right {
			mirror := 2*center - i
			p[i] = min(right-i, p[mirror])
		}

		// 中心扩展
		for t[i+(1+p[i])] == t[i-(1+p[i])] {
			p[i]++
		}

		// 更新最右边界和中心
		if i+p[i] > right {
			center, right = i, i+p[i]
		}

		// 更新最长回文
		if p[i] > maxLen {
			maxCenter, maxLen = i, p[i]
		}
	}

	// 还原原始字符串中的位置
	start := (maxCenter - maxLen) / 2
	return s[start : start+maxLen]
}

func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}
	result := make([]byte, 2*n+3)
	result[0] = '^'
	for i := 0; i < n; i++ {
		result[2*i+1] = '#'
		result[2*i+2] = s[i]
	}
	result[2*n+1] = '#'
	result[2*n+2] = '$'
	return string(result)
}
