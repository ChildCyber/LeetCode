package leetcode

// 字母大小写全排列
// https://leetcode.cn/problems/letter-case-permutation/
// 回溯
func letterCasePermutation(s string) []string {
	var ans []string
	backtrace([]byte(s), 0, &ans)
	return ans
}

func backtrace(str []byte, start int, ans *[]string) {
	*ans = append(*ans, string(str))
	for i := start; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' { // 小写字母
			str[i] -= 32 // 变为大写
			backtrace(str, i+1, ans)
			str[i] += 32 // 变为小写
		}
		if str[i] >= 'A' && str[i] <= 'Z' { // 大写字母
			str[i] += 32 // 变为小写
			backtrace(str, i+1, ans)
			str[i] -= 32 // 变为大写
		}
	}
}
