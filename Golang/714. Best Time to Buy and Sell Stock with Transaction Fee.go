package leetcode

// 买卖股票的最佳时机含手续费
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

// 动态规划
func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}

	// dp[i][0]: 第i天结束时，不持有股票的最大利润
	// dp[i][1]: 第i天结束时，持有股票的最大利润
	dp := make([][2]int, n)
	// 初始状态
	dp[0][0] = 0          // 第0天不持有股票，利润为0
	dp[0][1] = -prices[0] // 第0天持有股票，利润为-prices[0]

	for i := 1; i < n; i++ {
		// 第i天不持有股票：
		// 1. 前一天就不持有，今天什么也不做
		// 2. 前一天持有，今天卖出（要扣除手续费）
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)

		// 第i天持有股票：
		// 1. 前一天就持有，今天继续持有
		// 2. 前一天不持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	// 最后一天不持有股票的利润就是最大利润
	return dp[n-1][0]
}

func maxProfit714Optimized(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	hold := -prices[0] // 第一天买入
	cash := 0          // 第一天不买
	for i := 1; i < n; i++ {
		hold = max(hold, cash-prices[i])     // 要么延续，要么买入
		cash = max(cash, hold+prices[i]-fee) // 要么延续，要么卖出
	}
	return cash
}
