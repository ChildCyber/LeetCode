package leetcode

// 字符串中的第一个唯一字符
// https://leetcode.cn/problems/first-unique-character-in-a-string/
// 找到它的第一个不重复的字符，并返回它的索引

// 哈希表存储出现次数
func firstUniqChar(s string) int {
	// 对字符串进行两次遍历
	result := make([]int, 26)
	// 第一次遍历时，使用哈希表统计出字符串中每个字符出现的次数
	for i := 0; i < len(s); i++ {
		result[s[i]-'a']++
	}
	// 第二次遍历时，只要遍历到了一个只出现一次的字符，就返回它的索引
	for i := 0; i < len(s); i++ {
		if result[s[i]-'a'] == 1 { // 当前字符重复出现
			return i
		}
	}
	return -1
}
