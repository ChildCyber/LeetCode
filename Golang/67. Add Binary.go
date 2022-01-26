package leetcode

import "strconv"

// 二进制求和
// https://leetcode-cn.com/problems/add-binary/
func addBinary(a, b string) string { // 末尾对齐
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0') // 对应ascii码相减
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans // i对应的数字为 (carry+a[i]+b[i]) mod 2；int转str
		carry /= 2 // 下一位的进位为 (carry+a[i]+b[i])/2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
