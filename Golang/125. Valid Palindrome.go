package leetcode

import "strings"

// 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome/
func isPalindromeString(s string) bool {
	// 将空字符串定义为有效的回文串
	// 只考虑字母和数字字符，忽略字母的大小写
	s = strings.ToLower(s) // 全部转为小写
	// 双指针，每次比较指针指向的字符是否相等
	left, right := 0, len(s)-1
	for left < right {
		// 获取字符或数字的位置，不是的话继续向下找，直到下一个字符或数字
		for left < right && !isChar(s[left]) {
			left++
		}
		for left < right && !isChar(s[right]) {
			right--
		}
		// 比较左右指针指向的字符是否相等
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 判断c是否是字符或者数字
func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') { // 只判断小写和数字
		return true
	}
	return false
}
