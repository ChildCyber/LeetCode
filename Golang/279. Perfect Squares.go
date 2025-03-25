package leetcode

import "math"

// 完全平方数
// https://leetcode.cn/problems/perfect-squares/
// 完全平方数是一个整数，其值等于另一个整数的平方

// 动态规划
// 本质就是一个最少硬币数的问题，状态转移基本一致（dp[i] = min(dp[i - coin] + 1)）
func numSquares(n int) int {
	// 定义状态：dp[i]表示组成i所需的最少平方数
	// 状态转移：对于每个i，尝试所有可能的平方数 j²
	// 状态转移方程：dp[i] = min(dp[i], dp[i - j * j] + 1)
	//   i表示当前数字，j*j表示平方数，意思是要组成数字i，可以先组成数字i-j²（dp[i-j*j]），然后加上一个j²平方数（1）
	dp := make([]int, n+1)
	// 初始状态 dp[0]：数字0只需要0个完全平方数
	for i := 1; i <= n; i++ { // 从小到大地枚举i来计算dp[i]
		dp[i] = i                   // 最坏的情况，i全由1相加
		for j := 1; j*j <= i; j++ { // i的范围在0～√i之间
			dp[i] = min(dp[i], dp[i-j*j]+1) // dp[i-j*j]+1表示减去一个完全平方数j的完全平方后的数量加1就等于dp[i]
		}
	}

	return dp[n]
}

// 动态规划优化
func numSquares1(n int) int {
	dp := make([]int, n+1)

	// 初始化：最坏情况是用i个1相加
	for i := range dp {
		dp[i] = i
	}
	dp[0] = 0

	// 计算所有完全平方数
	maxSquare := int(math.Sqrt(float64(n)))
	squares := make([]int, maxSquare)
	for i := 1; i <= maxSquare; i++ {
		squares[i-1] = i * i
	}

	// 填充DP表
	for i := 1; i <= n; i++ {
		for _, square := range squares {
			if square > i {
				break // 平方数太大，跳过
			}
			// 状态转移
			dp[i] = min(dp[i], dp[i-square]+1)
		}
	}

	return dp[n]
}

// 数学-拉格朗日四平方定理
// 任何正整数都可以表示为不超过4个整数的平方和。也就是说答案只可能是1, 2, 3或4。
// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 勒让德三平方定理：一个数可以表示为三个平方数之和，当且仅当它不能表示为 4^a(8b+7) 的形式
// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}

func numSquaresMath(n int) int {
	// 检查是否是完全平方数
	if isPerfectSquare(n) {
		return 1
	}
	// 检查是否满足勒让德条件
	if checkAnswer4(n) {
		return 4
	}
	// 检查是否可以表示为两个平方数之和
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}
