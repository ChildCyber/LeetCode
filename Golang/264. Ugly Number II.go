package leetcode

// 丑数 II
// https://leetcode.cn/problems/ugly-number-ii/
// 动态规划
func nthUglyNumber(n int) int {
	// 定义数组 dp，其中 dp[i] 表示第 i 个丑数，第 n 个丑数即为 dp[n]。
	// 由于最小的丑数是 1，因此 dp[1]=1。
	// 定义三个指针 p2,p3,p5​，表示下一个丑数是当前指针指向的丑数乘以对应的质因数。初始时，三个指针的值都是 1
	dp, p2, p3, p5 := make([]int, n+1), 1, 1, 1
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(min(x2, x3), x5)
		// 分别比较 dp[i] 和 dp[p2]×2,dp[p3]×3,dp[p5]×5 是否相等，如果相等则将对应的指针加 111。
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}
