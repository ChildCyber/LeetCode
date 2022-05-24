package leetcode

import "math"

// 完全平方数
// https://leetcode.cn/problems/perfect-squares/
// 完全平方数是一个整数，其值等于另一个整数的平方，其值等于一个整数自乘的积

// 动态规划
func numSquares(n int) int {
	// dp[i]表示最少需要多少个数的平方来表示整数i
	// 状态转移方程：dp[i] = MIN(dp[i], dp[i - j * j] + 1)
	// i表示当前数字，j*j 表示平方数
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ { // 从小到大地枚举i来计算dp[i]
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}

// 数学
// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquaresMath(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
