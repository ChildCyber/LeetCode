package leetcode

// 回文数
// https://leetcode-cn.com/problems/palindrome-number
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// 反转数字
	var tmp, rev = x, 0
	for tmp != 0 {
		rev = rev*10 + tmp%10
		tmp /= 10
	}
	// 判断两数是否相等
	return rev == x
}
