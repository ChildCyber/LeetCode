package leetcode

// 买卖股票的最佳时机
// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/
// 只能买卖一次，在某日买入，在之后的某天卖出；lc 122可以多次买卖

// 暴力
func maxProfitBrute(prices []int) int {
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			profit := prices[j] - prices[i]
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

// 一次遍历
// 在历史最低点买入
func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}

	minPrice, maxProfit := prices[0], 0 // minPrice为历史最低买入价格，初始为第一天的价格；maxProfit为最大收益
	for i := 1; i < len(prices); i++ {  // 从第二天开始才可以卖出
		if prices[i]-minPrice > maxProfit { // 历史最低点买入，该日卖出能赚的钱
			maxProfit = prices[i] - minPrice // 最大收益=当天价格-历史最低买入价格
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return maxProfit
}

// 单调栈
func maxProfitStack(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	// 遍历每日价格，维护一个存储当日价格的单调栈
	// 栈底到栈顶对应的价格依次递增，栈底元素最小，栈顶元素最大
	stack, res := []int{prices[0]}, 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > stack[len(stack)-1] { // 当日价格大于栈顶价格，入栈
			stack = append(stack, prices[i])
		} else { // 当日价格小于栈顶价格，出栈
			index := len(stack) - 1
			for ; index >= 0; index-- { // 一直获取到小于该日价格的元素
				if stack[index] < prices[i] {
					break
				}
			}
			stack = stack[:index+1]
			stack = append(stack, prices[i])
		}
		res = max(res, stack[len(stack)-1]-stack[0])
	}
	return res
}
