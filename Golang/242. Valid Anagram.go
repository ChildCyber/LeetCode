package leetcode

import "sort"

// 有效的字母异位词
// https://leetcode-cn.com/problems/valid-anagram/

// 排序
func isAnagramSort(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	s1, s2 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })

	// 排序后的字符串是否相等
	return string(s1) == string(s2)
}

// 哈希表-Unicode版本
func isAnagramUnicode(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	cnt := map[rune]int{} // rune表示unicode

	// 统计s的字符频率
	for _, ch := range s {
		cnt[ch]++
	}
	// 递减并检查t的字符频率
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 { // 小于0
			return false
		}
	}

	// 由于已经检查了len(s)==len(t)，且在遍历t的过程中排除了t多余字符的情况，因此如果能完成循环，map中所有值必然为0（或已删除），返回true
	return true
}

// 数组替代哈希表
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := [26]int{}
	for i := 0; i < len(s); i++ {
		alphabet[s[i]-'a']++ // 遍历s，增加计数
		alphabet[t[i]-'a']-- // 遍历t，减少计数
	}

	// 判断值是否都是0
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 { // 有非0的数返回false
			return false
		}
	}

	return true
}

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
