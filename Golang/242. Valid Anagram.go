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
func isAnagram(s, t string) bool {
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

	// 判断值是否都是0
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 { // 有非0的数返回false
			return false
		}
	}
	return true
}

// 数组替代哈希表
func isAnagram1(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	// 判断数组是否相等
	return c1 == c2
}

func isAnagramUnicode(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{} // rune表示unicode
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 { // 小于0
			return false
		}
	}

	return true
}
