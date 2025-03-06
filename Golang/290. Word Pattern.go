package leetcode

import "strings"

// 单词规律
// https://leetcode.cn/problems/word-pattern/

// 哈希表-双向映射
func wordPattern(pattern string, s string) bool {
	// 分割字符串为单词切片
	words := strings.Split(s, " ")

	// 检查长度是否匹配
	if len(pattern) != len(words) {
		return false
	}

	// 创建两个映射表
	patternToWord := make(map[byte]string) // 字符->单词
	wordToPattern := make(map[string]byte) // 单词->字符

	for i := 0; i < len(pattern); i++ {
		char := pattern[i]
		word := words[i]

		// 检查pattern->word映射是否一致
		if mappedWord, ok := patternToWord[char]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			patternToWord[char] = word
		}

		// 检查word->pattern映射是否一致
		if mappedChar, ok := wordToPattern[word]; ok {
			if mappedChar != char {
				return false
			}
		} else {
			wordToPattern[word] = char
		}
	}

	return true
}
