package leetcode

import "strings"

// 翻转字符串里的单词
// https://leetcode-cn.com/problems/reverse-words-in-a-string
func reverseWords151(s string) string {
	ss := strings.Fields(s)       // 空格拆分字符串
	reverse151(&ss, 0, len(ss)-1) // 翻转
	return strings.Join(ss, " ")  // 连接
}

func reverse151(m *[]string, i int, j int) {
	for i <= j {
		(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
		i++
		j--
	}
}
