package leetcode

// 加一
// https://leetcode-cn.com/problems/plus-one/
// 找出最长的后缀 9
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			// no carry
			return digits
		}
		// carry
		digits[i] = 0
	}
	// digits 中所有的元素均为 9
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
