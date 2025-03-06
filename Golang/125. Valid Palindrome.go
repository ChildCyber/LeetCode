package leetcode

// 验证回文串
// https://leetcode-cn.com/problems/valid-palindrome/

// 双指针：原地比较
func isPalindromeString(s string) bool {
	// 将空字符串定义为有效的回文串，只考虑字母和数字字符，忽略字母的大小写
	// 双指针，每次比较指针指向的字符是否相等
	left, right := 0, len(s)-1
	for left < right {
		// 跳过非字母数字字符
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		// 比较字符（忽略大小写）
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}

		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c - 'A' + 'a'
	}
	return c
}
