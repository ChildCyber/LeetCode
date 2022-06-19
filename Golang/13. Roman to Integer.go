package leetcode

// 罗马数字转整数
// https://leetcode.cn/problems/roman-to-integer/

// 模拟
var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	num, lastInt, total := 0, 0, 0
	for i := 0; i < len(s); i++ {
		char := s[len(s)-1-i : len(s)-i] // 从后向前，s[len(s)-1-i]使用切片的形式避免类型转换
		num = roman[char]
		if num < lastInt { // 罗马数字中小的数字在大的数字的左边，需要减去小的数字
			total -= num
		} else { // 罗马数字中小的数字在大的数字的右边，可以将每个字符视作一个单独的值，累加每个字符对应的数值
			total += num
		}
		lastInt = num
	}
	return total
}

// 递归
var values = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToIntRec(s string) int {
	if len(s) == 1 {
		return values[s[0]]
	}
	// 处理s[1:]
	if values[s[0]] >= values[s[1]] { // 罗马数字中小的数字在大的数字的右边
		return romanToInt(s[1:]) + values[s[0]]
	}
	return romanToInt(s[1:]) - values[s[0]] // 罗马数字中小的数字在大的数字的左边
}
