package leetcode

// 单词拆分
// https://leetcode-cn.com/problems/word-break/

// 暴力（DFS）
func wordBreakBrute(s string, wordDict []string) bool {
	n := len(s)
	if n == 0 {
		return true // 空字符串可以被拆分
	}

	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	for i := 1; i <= n; i++ {
		if wordDictSet[s[:i]] { // 子字符串s[0:i]匹配
			if wordBreakBrute(s[i:n], wordDict) { // 递归判断s[i:n]
				return true
			}
		}
	}
	return false
}

// BFS
func wordBreakBFS(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 队列存储当前已经匹配到的位置
	queue := []int{0}                 // 从位置0开始
	visited := make([]bool, len(s)+1) // 避免重复访问

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]

		// 如果已经访问过这个位置，跳过
		if visited[start] {
			continue
		}
		visited[start] = true

		// 尝试所有可能的单词匹配
		for end := start + 1; end <= len(s); end++ {
			if wordSet[s[start:end]] {
				// 如果匹配到字符串末尾，返回成功
				if end == len(s) {
					return true
				}
				// 将新的起始位置加入队列
				queue = append(queue, end)
			}
		}
	}

	return false
}

// 动态规划-不带剪枝
// 时间复杂度：O(n²)
// 空间复杂度：O(n+m)
func wordBreak(s string, wordDict []string) bool {
	// 将字典转换为哈希集合，提高查找效率
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}

	// 状态定义：dp[i] 表示字符串 s 前 i 个字符组成的字符串，即 s[0:i] 是否能被空格拆分成字典中的单词
	// 状态转移方程：对于每个位置 i，检查所有可能的分割点 j（0 ≤ j < i）：
	//   如果 dp[j] 为真（前j个字符可拆分）并且子串 s[j:i] 在字典中，那么 dp[i] 也为真
	dp := make([]bool, len(s)+1)
	dp[0] = true                   // 空字符串可以被拆分
	for i := 1; i <= len(s); i++ { // 遍历字符串的所有子串
		for j := 0; j < i; j++ { // 得到字符串s[j:i]
			if dp[j] && wordDictSet[s[j:i]] { // s[0:j]可以被匹配且s[j:i]在在字典中，s[0:i]可以被匹配
				dp[i] = true // s的前i位可以用wordDict中的单词表示
				break        // 只匹配最长的子字符串，0<=j<i，找到一个可行的拆分方式即可
			}
		}
	}

	// 整个字符串能否被拆分
	return dp[len(s)]
}

// 限制检查范围
func wordBreakOptimized(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	maxWordLen := 0 // 记录字典中最长单词的长度

	for _, word := range wordDict {
		wordSet[word] = true
		maxWordLen = max(len(word), maxWordLen)
	}

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 1; i <= n; i++ {
		// 优化：只检查可能成为单词的长度范围
		start := max(0, i-maxWordLen)
		for j := start; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[n]
}

// 记忆化递归（DFS + 记忆化）
// 自顶向下递归
func wordBreakTop(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 记忆化缓存，memo[i]表示从位置i开始的子串能否被拆分
	memo := make(map[int]bool)
	return canBreak(s, 0, wordSet, memo)
}

func canBreak(s string, start int, wordSet map[string]bool, memo map[int]bool) bool {
	// 基准情况：已经到达字符串末尾
	if start == len(s) {
		return true
	}

	// 如果已经计算过这个位置，直接返回结果
	if val, exists := memo[start]; exists {
		return val
	}

	// 尝试所有可能的结束位置
	for end := start + 1; end <= len(s); end++ {
		// 如果当前子串在字典中，且剩余部分也能被拆分
		if wordSet[s[start:end]] && canBreak(s, end, wordSet, memo) {
			memo[start] = true // 缓存结果
			return true
		}
	}

	memo[start] = false // 缓存结果
	return false
}
