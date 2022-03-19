package leetcode

// Excel 表列序号
// https://leetcode.cn/problems/excel-sheet-column-number/
func titleToNumber(s string) int {
	val, ans := 0, 0
	for i := 0; i < len(s); i++ {
		val = int(s[i] - 'A' + 1)
		ans = ans*26 + val
	}
	return ans
}
