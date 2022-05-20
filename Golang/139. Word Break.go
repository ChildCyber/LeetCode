package leetcode

// 单词拆分
// https://leetcode-cn.com/problems/word-break/

// 暴力
func wordBreakForce(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true
	}

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	for i := 1; i <= n; i++ {
		if wordDictSet[s[:i]] {
			if wordBreakForce(s[i:n], wordDict) { // 递归判断s[i:n]
				return true
			}
		}
	}
	return false
}

// 动态规划
// 不带剪枝
func wordBreak(s string, wordDict []string) bool {
	// 检查一个字符串是否出现在给定的字符串列表中，slice变map，方便查询
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	// dp[i] 表示字符串 s 前 i 个字符组成的字符串 s[0..i−1] 是否能被空格拆分成若干个字典中出现的单词
	dp := make([]bool, len(s)+1)
	dp[0] = true                   // 空字符可以被表示
	for i := 1; i <= len(s); i++ { // 遍历字符串的所有子串
		for j := 0; j < i; j++ { // 得到字符串s[j:i]
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true // s的前i位是否可以用wordDict中的单词表示
				break
			}
		}
	}

	return dp[len(s)]
}
