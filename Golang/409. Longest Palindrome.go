package leetcode

// 最长回文串
// https://leetcode-cn.com/problems/longest-palindrome/
// 给定一个包含大写字母和小写字母的字符串 s，返回 通过这些字母构造成的 最长的回文串
// 贪心
// 分为：奇数个字符回文串、偶数个字符回文串
func longestPalindrome409(s string) int { // 返回长度
	// 创建哈希表，统计每个字符的频次
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	ans := 0 // 记录最长回文串长度
	// 使所有字符尽可能多的出现偶数 + 1次（如果此时）
	for _, v := range counter { // key, value
		ans += v / 2 * 2            // v 变为偶数
		if ans%2 == 0 && v%2 == 1 { // 该字符出现奇数次且ans为偶数
			// 在发现了第一个出现次数为奇数的字符后，将 ans 增加 1，这样 ans 变为奇数，在后面发现其它出现奇数次的字符时，就不改变 ans 的值了
			ans++
		}
	}
	return ans
}
