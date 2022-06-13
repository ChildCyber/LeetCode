package leetcode

// 加一
// https://leetcode-cn.com/problems/plus-one/
// 找出最长的后缀 9
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{}
	}

	carry := 1
	for i := len(digits) - 1; i >= 0; i++ { // 逆序遍历，找出第一个不为 9 的元素，将其加一并将后续所有元素置零即可
		if digits[i]+carry > 9 {
			digits[i] = 0
			carry = 1
		} else {
			digits[i] += carry
			carry = 0
		}
	}

	if digits[0] == 0 && carry == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
