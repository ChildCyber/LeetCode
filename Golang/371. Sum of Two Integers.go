package leetcode

// 两整数之和
// https://leetcode.cn/problems/sum-of-two-integers/
// 位运算
func getSum(a, b int) int {
	for b != 0 {
		carry := uint(a&b) << 1 // 进位
		a ^= b                  // 不带进位的加法
		b = int(carry)
	}
	return a
}

func getSum1(a int, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	// (a & b)<<1 计算的是进位
	// a ^ b 计算的是不带进位的加法
	return getSum1((a&b)<<1, a^b)
}
