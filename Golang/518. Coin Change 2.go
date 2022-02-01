package leetcode

// 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/
// 动态规划
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	// 组合问题，不关心硬币使用的顺序，而是硬币有没有被用到
	for _, coin := range coins { // 枚举硬币
		for i := coin; i <= amount; i++ { // 枚举金额
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
