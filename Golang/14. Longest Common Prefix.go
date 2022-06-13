package leetcode

// 最长公共前缀
// https://leetcode-cn.com/problems/longest-common-prefix

// 纵向扫描
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 从前往后遍历所有字符串的每一列，比较相同列上的字符是否相同，如果相同则继续对下一列进行比较，如果不相同则当前列不再属于公共前缀，当前列之前的部分为最长公共前缀
	for i := 0; i < len(strs[0]); i++ { // 字符数组中首个字符串
		for j := 1; j < len(strs); j++ { // 除了第一个字符串外，字符数组中所有字符串
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}

	return strs[0]
}
