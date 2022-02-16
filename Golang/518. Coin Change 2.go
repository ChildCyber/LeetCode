package leetcode

// 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/
// 动态规划
func change(amount int, coins []int) int { // 返回可以凑成总金额的硬币组合数
	// 状态转移方程： dp[i] += dp[i-coin]，coin 为当前枚举的 coin
	// 用 dp[x] 表示金额之和等于 x 的硬币组合数，目标是求 dp[amount]
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
