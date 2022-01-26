package leetcode

// 最后一个单词的长度
// https://leetcode-cn.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 { // 先从第一个没有空格的单词的最后一个字母起
				continue
			}
			break
		}
		count++
	}
	return count
}
