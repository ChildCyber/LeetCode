package leetcode

import "strings"

// 找出字符串中第一个匹配项的下标
// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/

// 内置函数
func strStrBuiltIn(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 暴力
func strStr(haystack, needle string) int {
	m, n := len(haystack), len(needle)

	if n == 0 {
		return 0
	}
	if m < n {
		return -1
	}

	for i := 0; i <= m-n; i++ {
		if haystack[i:i+n] == needle {
			return i
		}
	}

	return -1
}
