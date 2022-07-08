package leetcode

import "math"

// 字符串转换整数 (atoi)
// https://leetcode.cn/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	abs, sign, i, n := 0, 1, 0, len(s) // ans,正负号,string索引,string长度
	// 丢弃无用的前导空格
	for i < n && s[i] == ' ' {
		i++
	}

	// 标记正负号
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	// 将数字转换成数字类型
	for i < n && s[i] >= '0' && s[i] <= '9' {
		abs = 10*abs + int(s[i]-'0') // 字节 byte '0' == 48

		// 判断是否超过 int 类型的上限 [-2^31, 2^31 - 1]，如果超过上限，需要输出边界，即-2^31 或者 2^31 - 1
		if sign*abs < math.MinInt32 {
			return math.MinInt32
		} else if sign*abs > math.MaxInt32 {
			return math.MaxInt32
		}

		i++
	}

	return sign * abs
}
