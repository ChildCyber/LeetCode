package leetcode

// 整数反转
// https://leetcode-cn.com/problems/reverse-integer
func reverse7(x int) int {
	var tmp int
	for x != 0 {
		tmp = tmp*10 + x%10
		x /= 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}
