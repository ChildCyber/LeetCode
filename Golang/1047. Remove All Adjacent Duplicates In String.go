package leetcode

// 删除字符串中的所有相邻重复项
// https://leetcode-cn.com/problems/remove-all-adjacent-duplicates-in-string
// 栈
func removeDuplicates1047(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
