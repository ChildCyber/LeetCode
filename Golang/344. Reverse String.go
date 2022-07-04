package leetcode

// 反转字符串
// https://leetcode.cn/problems/reverse-string/
// 双指针
func reverseString(s []byte) {
	for i, n := 0, len(s); i < n/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}
