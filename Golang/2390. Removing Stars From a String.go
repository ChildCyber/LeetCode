package leetcode

// 从字符串中移除星号
// https://leetcode.cn/problems/removing-stars-from-a-string/

// 栈
func removeStars(s string) string {
	ans := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			ans = ans[:len(ans)-1]
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}
