package leetcode

import "math"

// 最佳买卖股票时机含冷冻期
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/
// 动态规划
func maxProfit309(prices []int) int {
	// 冷冻期为1天：卖出股票后，无法在第二天买入股票
	length := len(prices)
	if length <= 1 {
		return 0
	}

	// price[i-1] 代表第i天的股票价格(由于 price 是从 0 开始的)
	// buy[i] 代表第i天通过 buy 或者 冷冻期 结束此天能获得的最大收益
	// sell[i] 代表第i天通过 sell 或者 冷冻期 结束此天能获得的最大收益
	buy, sell := make([]int, length), make([]int, length)
	for i := range buy {
		buy[i] = math.MinInt64
	}
	buy[0] = -prices[0]                      // 代表第一天就买入，所以收益变成负
	buy[1] = max(buy[0], -prices[1])         // 代表第一天不买入，第二天再买入
	sell[0] = 0                              // 代表第一天卖出，由于第一天不持有股票，所以 sell[0] = 0
	sell[1] = max(sell[0], buy[0]+prices[1]) // 代表第一天卖出和第一天买入、第二天卖做对比

	// 状态转移：
	// 在不违反规则的前提下进行「买入」或者「卖出」操作，此时第 i 天的状态会从第 i−1 天的状态转移而来
	// 也可以不进行任何操作，此时第 i 天的状态就等同于第 i−1 天的状态
	for i := 2; i < length; i++ {
		sell[i] = max(sell[i-1] /*冷冻期*/, buy[i-1]+prices[i] /*sell*/)
		buy[i] = max(buy[i-1] /*冷冻期*/, sell[i-2]-prices[i] /*buy*/) // 第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1
	}

	return sell[len(sell)-1]
}
