package leetcode

// 解码方法
// https://leetcode-cn.com/problems/decode-ways/
// 动态规划
func numDecodings(s string) int {
	n := len(s)
	// dp[i] 表示字符串 s 的前 i 个字符 s[1..i] 的解码方法数
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && (s[i-2]-'0')*10+(s[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
