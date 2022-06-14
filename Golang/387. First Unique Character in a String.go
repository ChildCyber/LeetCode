package leetcode

// 字符串中的第一个唯一字符
// https://leetcode.cn/problems/first-unique-character-in-a-string/
func firstUniqChar(s string) int {
	result := make([]int, 26)
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 { // 当前字符重复出现
			return i
		}
	}
	return -1
}
