package leetcode

// 字符串的最大公因子
// https://leetcode.cn/problems/greatest-common-divisor-of-strings/

func gcdOfStrings(str1, str2 string) string {
	// 必要条件：若存在公因子 x，使得 str1 = x^a, str2 = x^b，
	// 则必有 str1+str2 == str2+str1（相当于 x^(a+b)）
	if str1+str2 != str2+str1 {
		return ""
	}

	g := gcd(len(str1), len(str2))
	if g == 0 { // 处理两者都为空的情况
		return ""
	}
	// 取长度为 gcd 的前缀
	if len(str1) >= g {
		return str1[:g]
	}
	return str2[:g]
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
