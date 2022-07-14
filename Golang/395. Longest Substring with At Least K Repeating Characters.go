package leetcode

import "strings"

// 至少有 K 个重复字符的最长子串
// https://leetcode.cn/problems/longest-substring-with-at-least-k-repeating-characters/
// 分治
func longestSubstring(s string, k int) (ans int) { // 子字符串中每个字符出现的次数都最少为k
	if s == "" {
		return
	}

	cnt := [26]int{}
	for _, ch := range s { // 统计每个字符出现次数
		cnt[ch-'a']++
	}

	var split byte
	for i, c := range cnt[:] { // 遍历cnt
		if 0 < c && c < k { // 如果某个字符出现次数大于 0 小于 k，那么包含这个字符的子串都不满足要求
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 { // 字符出现次数大于k，返回
		return len(s)
	}

	for _, subStr := range strings.Split(s, string(split)) { // 将字符串切分成多个小段分别求解
		ans = max(ans, longestSubstring(subStr, k))
	}

	return ans
}
