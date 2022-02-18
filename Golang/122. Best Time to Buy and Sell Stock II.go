package leetcode

// 买卖股票的最佳时机 II
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/
// 不止买卖一次，可以买卖多次，但买卖不能在同一天内操作

// 贪心：最大收益来源就是每次跌了就买入，涨到顶峰的时候就抛出
func maxProfit122(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ { // 小于len(prices)-1，[0:len(prices)-2]
		if prices[i+1] > prices[i] { // 只要有涨峰就开始计算赚的钱，连续涨可以用两两相减累加来计算，两两相减累加，相当于涨到波峰的最大值减去谷底的值
			profit += prices[i+1] - prices[i]
		}
	}
	return profit
}

// 动态规划
func maxProfit122DP(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
