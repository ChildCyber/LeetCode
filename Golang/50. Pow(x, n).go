package leetcode

// Pow(x, n)
// https://leetcode-cn.com/problems/powx-n
// 快速幂 + 递归
func myPow(x float64, n int) float64 {
	// 指数n为正数
	if n >= 0 {
		return quickMul(x, n)
	}
	// 指数n为负数
	return 1.0 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	// 计算 x^n 时，先递归地计算出 x^⌊n/2⌋
	y := quickMul(x, n/2)
	// n为偶数
	if n%2 == 0 {
		return y * y
	}
	// n为奇数
	return y * y * x
}

// 迭代
func quickMul1(x float64, N int) float64 {
	ans := 1.0
	// 贡献的初始值为 x
	xContribute := x
	// 在对 N 进行二进制拆分的同时计算答案
	for N > 0 {
		if N%2 == 1 {
			// 如果 N 二进制表示的最低位为 1，那么需要计入贡献
			ans *= xContribute
		}
		// 将贡献不断地平方
		xContribute *= xContribute
		// 舍弃 N 二进制表示的最低位，这样我们每次只要判断最低位即可
		N /= 2
	}
	return ans
}
