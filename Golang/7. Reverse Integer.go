package leetcode

// 整数反转
// https://leetcode-cn.com/problems/reverse-integer
func reverse7(x int) int {
	var ans int
	for x != 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	// 范围在[−2^31, 2^31−1]
	if ans > 1<<31-1 || ans < -(1<<31) {
		return 0
	}
	return ans
}
