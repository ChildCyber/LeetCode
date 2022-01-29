package leetcode

import "strconv"

// 字符串相加
// https://leetcode-cn.com/problems/add-strings/
// 模拟
func addStrings(num1 string, num2 string) string {
	add := 0
	ans := ""
	// 定义两个指针 i 和 j 分别指向 num1​ 和 num2 的末尾，即最低位，同时定义一个变量 add 维护当前是否有进位，然后从末尾到开头逐位相加即可
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans // 直接字符拼接
		add = result / 10
	}
	return ans
}
