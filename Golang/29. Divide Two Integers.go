package leetcode

import "math"

// 两数相除
// https://leetcode.cn/problems/divide-two-integers/

// 暴力
func divideForce(dividend, divisor int) int {
	var sign, ans int
	if (dividend^divisor)>>31&1 == 1 { // 除数被除数符号不同
		sign = -1
	} else { // 符号相同
		sign = 1
	}

	// 根据除法定义，不停减去除数
	for dividend >= divisor {
		dividend -= divisor
		ans++
	}

	ans *= sign
	if ans > math.MaxInt64 || ans < math.MinInt64 {
		return math.MaxInt64
	}

	return ans
}

// 指数递增步长
func divide(dividend, divisor int) int {
	var sign, ans int
	if (dividend^divisor)>>31&1 == 1 { // 除数被除数符号不同
		sign = -1
	} else { // 符号相同
		sign = 1
	}

	// 根据除法定义，不停减去除数
	// 翻倍步长快速逼近，每次步长加倍
	for dividend >= divisor {
		i := 1
		tmp := divisor
		for dividend >= tmp {
			dividend -= tmp
			ans += i
			i <<= 1
			tmp <<= 1
		}
	}

	ans *= sign
	if ans > math.MaxInt64 || ans < math.MinInt64 {
		return math.MaxInt64
	}

	return ans
}
