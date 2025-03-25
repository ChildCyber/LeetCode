package leetcode

// 罗马数字转整数
// https://leetcode.cn/problems/roman-to-integer/

// 问题本质：
// 罗马数字中，小数字在大数字前表示减法（如IV=4），否则表示加法（如VI=6）
// 关键是判断当前字符与下一个字符的大小关系

// 模拟
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	total := 0
	num, lastInt := 0, 0 // 当前数字，前一个数字
	for i := 0; i < len(s); i++ {
		char := s[len(s)-1-i : len(s)-i] // 字符s[len(s)-1-i]，使用切片的形式避免类型转换
		num = romanMap[char]
		if num < lastInt { // 罗马数字中小的数字在大的数字的左边，需要减去小的数字
			total -= num
		} else { // 罗马数字中小的数字在大的数字的右边，可以将每个字符视作一个单独的值，累加每个字符对应的数值
			total += num
		}
	}
	return total
}

// 递归
func romanToIntRecursive(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	if len(s) == 1 {
		return romanMap[s[0]]
	}

	// 处理s[1:]
	if romanMap[s[0]] >= romanMap[s[1]] { // 罗马数字中小的数字在大的数字的右边
		return romanToIntRecursive(s[1:]) + romanMap[s[0]]
	}

	return romanToIntRecursive(s[1:]) - romanMap[s[0]] // 罗马数字中小的数字在大的数字的左边
}
