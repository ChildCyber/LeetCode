package leetcode

// 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 只能买卖一次，在某日买入，在之后的某天卖出；122可以多次买卖
func maxProfit(prices []int) int { //dp
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0     // min为历史最低买入价格，初始为第一天的价格；maxProfit为最大收益
	for i := 1; i < len(prices); i++ { // 从第二天开始才可以卖出
		// price[i]为目前为止最高价格
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min // 最大收益=当天价格-历史最低买入价格
		}
		// 记录目前为止，历史最低买入价格
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}

// 单调栈
