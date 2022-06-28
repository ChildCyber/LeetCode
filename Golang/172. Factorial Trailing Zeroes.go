package leetcode

// 阶乘后的零
// https://leetcode.cn/problems/factorial-trailing-zeroes/
// 数学
// 有多少个5就有多少个0
func trailingZeroes(n int) (ans int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return
}
