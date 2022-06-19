package leetcode

import "sort"

// 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/
// 排序
func isAnagramSort(s, t string) bool {
	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	// 排序后的字符串是否相等
	return string(s1) == string(s2)
}

// 哈希表
func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}

	for i := 0; i < len(sBytes); i++ { // 记录字符出现频次
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ { // 减去字符出现频次
		alphabet[tBytes[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}
