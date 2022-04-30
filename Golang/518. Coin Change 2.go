package leetcode

// 零钱兑换 II
// https://leetcode-cn.com/problems/coin-change-2/
// 返回可以凑成总金额的硬币组合数

// 暴力递归
func changeForce(amount int, coins []int) int {
	var dfs func([]int, int, int, int) int
	dfs = func(coins []int, amount int, index int, sum int) int {
		if amount == sum {
			return 1
		}
		if sum > amount || index >= len(coins) {
			return 0
		}

		ans := 0
		// 选择当前硬币，并且下一个继续可以选择当前硬币
		ans += dfs(coins, amount, index, sum+coins[index])
		// 不选当前硬币，并且跳过当前硬币
		ans += dfs(coins, amount, index+1, sum)
		return ans
	}
	return dfs(coins, amount, 0, 0)
}

// 动态规划
func change(amount int, coins []int) int {
	// 状态转移方程： dp[i] += dp[i-coin]，coin 为当前枚举的 coin
	// 用 dp[x] 表示金额之和等于 x 的硬币组合数，目标是求 dp[amount]
	dp := make([]int, amount+1)
	dp[0] = 1

	// coins 的每个元素可以选取多次，且不考虑选取元素的顺序
	// 对于面额为 coin 的硬币，当 coin≤i≤amount 时，如果存在一种硬币组合的金额之和等于 i−coin，则在该硬币组合中增加一个面额为 coin 的硬币，即可得到一种金额之和等于 i 的硬币组合
	for _, coin := range coins { // 枚举硬币
		for i := coin; i <= amount; i++ { // 枚举金额，coin≤i≤amount，由于顺序确定，因此不会重复计算不同的排列
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}
