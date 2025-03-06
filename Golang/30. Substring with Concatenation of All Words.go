package leetcode

// 串联所有单词的子串
// https://leetcode.cn/problems/substring-with-concatenation-of-all-words/

// 暴力
// 枚举所有可能的子串起始位置，检查是否包含所有单词

// 滑动窗口
func findSubstring(s string, words []string) []int {
	ans := []int{}
	if len(words) == 0 || len(s) == 0 {
		return ans
	}

	wordLen := len(words[0])
	wordCount := len(words)
	totalLen := wordLen * wordCount

	if len(s) < totalLen {
		return ans
	}

	// 统计目标词频
	need := make(map[string]int)
	for _, w := range words {
		need[w]++
	}

	// 遍历所有偏移量（所有可能的起始位置）
	for i := 0; i < wordLen; i++ {
		left := i
		count := 0
		window := make(map[string]int)

		// 按单词步进（这里开始是滑动窗口的匹配部分）
		for j := i; j+wordLen <= len(s); j += wordLen {
			word := s[j : j+wordLen] // 窗口内匹配的单词
			if _, ok := need[word]; ok {
				window[word]++
				count++

				// 超出需要，左边收缩
				for window[word] > need[word] {
					leftWord := s[left : left+wordLen] // 移除左边界单词
					window[leftWord]--
					left += wordLen
					count--
				}

				// 满足条件
				if count == wordCount {
					ans = append(ans, left)
					// 移动左边界继续寻找
					leftWord := s[left : left+wordLen]
					window[leftWord]--
					left += wordLen
					count--
				}
			} else {
				// 遇到不在need中的单词，重置窗口
				window = make(map[string]int)
				count = 0
				left = j + wordLen
			}
		}
	}

	return ans
}
