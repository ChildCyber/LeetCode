package leetcode

// 3 的幂
// https://leetcode.cn/problems/power-of-three/
// 试除法
func isPowerOfThree(n int) bool {
	for n > 0 && n%3 == 0 {
		n /= 3
	}
	return n == 1
}
