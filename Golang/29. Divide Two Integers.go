package leetcode

import "math"

// 两数相除
// https://leetcode.cn/problems/divide-two-integers/

// 减法代替除法，暴力
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func divideForce(dividend, divisor int) int {
	var sign, ans int // 1为正数，-1为负数

	if (dividend^divisor)>>31&1 == 1 { // 除数被除数符号不同
		sign = -1
	} else { // 符号相同
		sign = 1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

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

// 减法代替除法，指数递增步长
// 时间复杂度：O(log n * log n)
// 空间复杂度：O(1)
func divide(dividend, divisor int) int {
	var sign, ans int // 1为正数，-1为负数

	if (dividend^divisor)>>31&1 == 1 { // 除数被除数符号不同
		sign = -1
	} else { // 符号相同
		sign = 1
	}

	dividend = abs(dividend)
	divisor = abs(divisor)

	// 根据除法定义，被除数不停减去除数，每次尝试减去除数的倍数
	for dividend >= divisor {
		i := 1
		tmp := divisor
		for dividend >= tmp {
			dividend -= tmp
			ans += i
			i <<= 1   // i*=2
			tmp <<= 1 // 每次步长加倍，tmp*=2
		}
	}

	ans *= sign
	// 32 位有符号整数
	if ans > math.MaxInt64 || ans < math.MinInt64 {
		return math.MaxInt64
	}

	return ans
}
