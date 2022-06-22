package leetcode

// 解码方法
// https://leetcode-cn.com/problems/decode-ways/
// 动态规划
func numDecodings(s string) int { // 返回解码方法的总数
	n := len(s)
	// dp[i] 表示字符串 s 的前 i 个字符 s[1..i] 的解码方法数
	dp := make([]int, n+1)
	dp[0] = 1 // 空字符串可以有 1 种解码方法，解码出一个空字符串

	for i := 1; i <= n; i++ {
		// 使用了一个字符， s[i] 进行解码
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		// 使用了两个字符， s[i−1] 和 s[i] 进行编码
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 { // s[i−1] 不能等于 0，且 s[i−1] 和 s[i] 组成的整数必须小于等于 26
			dp[i] += dp[i-2]
		}
	}

	return dp[n]
}
