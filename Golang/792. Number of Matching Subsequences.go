package leetcode

// 匹配子序列的单词数
// https://leetcode.cn/problems/number-of-matching-subsequences/
// 暴力
func numMatchingSubseqForce(s string, words []string) int {
	ans := 0
	for _, word := range words {
		if len(words) > len(s) {
			continue
		}

		if isSubSeq(s, word) {
			ans++
		}
	}
	return ans
}

// 双指针，判断w是否为s的子序列
func isSubSeq(s, w string) bool {
	i, j := 0, 0
	for ; i < len(s) && j < len(w); i++ { // 遍历s
		if s[i] == w[j] {
			j++
		}
		if j == len(w) {
			return true
		}
	}
	return false
}

// 哈希表
func numMatchingSubseq(s string, words []string) int {
	m := map[string]bool{}
	ans := 0

	for _, w := range words {
		if m[w] { // 减少重复字符串判断
			ans++
			continue
		}

		isSub := isSubSeq(s, w)
		m[w] = isSub
		if isSub {
			ans++
		}

	}
	return ans
}
