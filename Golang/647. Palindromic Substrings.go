package leetcode

// 回文子串
// https://leetcode.cn/problems/palindromic-substrings/

// 暴力
// 枚举出所有的子串，然后再判断这些子串是否是回文
// 时间复杂度：O(n^3)
func countSubstringsBrute(s string) int {
	n := len(s)
	if n < 1 {
		return 0
	}
	ans := 0

	var isPalindrome func(i, j int) bool
	isPalindrome = func(i, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}

	for i := 0; i < n; i++ { // 分割后字符串起始字符
		for j := i; j < n; j++ { // 分割后字符串末尾字符
			if isPalindrome(i, j) { // 判断子串是否是回文
				ans++
			}
		}
	}
	return ans
}

// 中心扩散法
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	// 从左往右扫描字符串，以每个字符做轴用中心扩散法，依次遍历计数回文子串
	for i := 0; i < 2*n-1; i++ { // 要考虑奇数和偶数两种回文：一个字符串长度是n，中心可能的个数 = n（奇数中心）+ n-1（偶数中心）= 2n-1，所以循环 i 从 0 到 2*n-2（即 <2*n-1）遍历所有中心
		// 重点：用 i%2 判断中心是奇数还是偶数
		// i%2 == 0：奇数长度回文，左右指针重合在一个字符上
		// i%2 == 1：偶数长度回文，左右指针在相邻字符之间
		l, r := i/2, i/2+i%2
		for 0 <= l && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
