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
		//char := s[len(s)-(i+1) : len(s)-i]
		char := s[len(s)-1-i : len(s)-i]
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
func romanToIntRec(s string) int {
	var values = map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	if len(s) == 1 {
		return values[s[0]]
	}
	if values[s[0]] >= values[s[1]] {
		return romanToInt(s[1:]) + values[s[0]]
	}
	return romanToInt(s[1:]) - values[s[0]]
}
