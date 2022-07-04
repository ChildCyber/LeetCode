package leetcode

import "strings"

// 实现strStr()
// https://leetcode-cn.com/problems/implement-strstr

// 内置函数
func strStrBuiltIn(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 暴力匹配
func strStr(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	m, n := len(haystack), len(needle)
	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
