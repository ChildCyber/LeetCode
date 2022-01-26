package leetcode

// 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix
func longestCommonPrefix(strs []string) string { // 纵向扫描
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ { // 字符数组中首个字符串
		for j := 1; j < len(strs); j++ { // 除了第一个字符串外，字符数组中所有字符串
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
