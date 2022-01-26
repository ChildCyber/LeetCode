package leetcode

// 最长回文串

// 贪心
// 奇数个字符回文串、偶数个字符回文串
// 先统计每个字符的频次,然后每个字符能取 2 个的取 2 个,不足 2 个的并且当前构造中的回文串是偶数的情况下(即每 2 个都配对了),可以取 1 个。最后组合出来的就是最⻓回文串。
func longestPalindrome(s string) int {
	// 创建哈希表
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	ans := 0
	for _, v := range counter { // key, value
		ans += v / 2 * 2            // v 变为偶数
		if ans%2 == 0 && v%2 == 1 { // 该字符出现奇数次且ans为偶数
			ans++
		}
	}
	return ans
}
