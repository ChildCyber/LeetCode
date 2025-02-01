package leetcode

import "sort"

// 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/

// 使用排序和计数作为哈希表的键

// 排序
// 时间复杂度：O(nk*logk)
func groupAnagrams(strs []string) [][]string {
	// key是排序以后的字符串，value对应的是排序字符串的字符串集合
	mp := map[string][]string{}
	for _, str := range strs { // 对每个字符串排序，排序之后得到的字符串一定是相同的
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		mp[key] = append(mp[key], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

// 计数
// 时间复杂度：O(nk)
func groupAnagramsCount(strs []string) [][]string {
	// 互为字母异位词的两个字符串包含的字母相同，因此两个字符串中的相同字母出现的次数一定相同
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}

	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
