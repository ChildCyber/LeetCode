package leetcode

// 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 不止买卖一次，可以买卖多次，但买卖不能在同一天内操作

// 贪心
// 思路：只要今天价格比昨天高，就完成一次交易（昨天买今天卖）
func maxProfitII(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// 动态规划
func maxProfitIIDP(prices []int) int {
	n := len(prices)
	// 状态定义：
	//   dp[i][0]：第 i 天结束时，不持有股票的最大利润
	//   dp[i][1]：第 i 天结束时，持有股票的最大利润
	// 状态转移方程：
	// 第i天不持有股票的最大利润 = max(昨天就不持有, 昨天持有今天卖出)
	//   dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// 第i天持有股票的最大利润 = max(昨天就持有, 昨天不持有今天买入)
	//   dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
