package leetcode

// 验证回文串 II
// https://leetcode.cn/problems/valid-palindrome-ii/

// 贪心+双指针
func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	// 使用左右指针从两端向中间移动，检查是否相等
	for left < right {
		// 思维转换：把"删除字符"转化为"检查子串是否是严格回文"
		if s[left] != s[right] {
			// 遇到第一个不匹配，尝试两种删除策略（不继续往后检查）：
			// 删除左指针的字符（检查s[left+1:right+1]）
			// 删除右指针的字符（检查s[left:right]）
			// 只要有一种策略能形成回文就返回true
			return isPalindromeStr(s, left+1, right) || isPalindromeStr(s, left, right-1)
		}
		left++
		right--
	}
	return true
}

// 检查子串s[start:end+1]是否是回文
func isPalindromeStr(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
