package leetcode

import "math"

// 零钱兑换
// https://leetcode-cn.com/problems/coin-change/
// 递归
func coinChangeRec(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	ans := math.MaxInt32

	var findWay func([]int, int, int)
	findWay = func(coins []int, amount int, count int) {
		if amount < 0 {
			return
		}
		if amount == 0 {
			ans = min(ans, count)
		}
		for i := 0; i < len(coins); i++ {
			findWay(coins, amount-coins[i], count+1)
		}
	}

	findWay(coins, amount, 0)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 记忆化搜索
func coinChange1(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	memo := make([]int, amount)
	// memo[n] 表示钱币n可以被换取的最少的硬币数，不能换取就为-1
	// findWay函数的目的是为了找到 amount数量的零钱可以兑换的最少硬币数量，返回其值int
	var findWay func([]int, int) int
	findWay = func(coins []int, amount int) int {
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}
		// 记忆化处理，memo[n]用赋予了值，就不用继续下面的循环
		// 直接的返回memo[n] 的最优值
		if memo[amount-1] != 0 {
			return memo[amount-1]
		}
		minVal := math.MaxInt32
		for i := 0; i < len(coins); i++ {
			ans := findWay(coins, amount-coins[i])
			if ans >= 0 && ans < minVal {
				minVal = ans + 1 // 加1，是为了加上得到res结果的那个步骤中，兑换的一个硬币
			}
		}
		tmp := -1
		if minVal == math.MaxInt32 {
			tmp = -1
		} else {
			tmp = minVal
		}
		memo[amount-1] = tmp
		return memo[amount-1]
	}
	return findWay(coins, amount)
}

// 二维DP
func coinChange2(coins []int, amount int) int {
	n := len(coins)
	// dp[i][j] 表示从[0,i]个硬币中需要选取硬币总和为j的个数
	// 默认值 -1 表示不能从[0,i]个硬币中凑成总和为j
	dp := [][]int{}
	for i := 0; i < n; i++ {
		tmp := []int{}
		for j := 0; j < amount+1; j++ {
			if j == 0 { // 金额为0
				tmp = append(tmp, 0)
			} else {
				tmp = append(tmp, -1)
			}
		}
		dp = append(dp, tmp)
	}

	for i := 1; i < n; i++ { // 硬币
		for j := 1; j < amount+1; j++ { // 金额
			coin := coins[i-1]
			// 如果硬币面值大于amount，或者是当前amount不能由上一个j - coin得到，则从上一个状态转移
			if coin > j || dp[i][j-coin] == -1 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 如果dp[i - 1][j] != -1说明上一个状态可以凑成amount
				if dp[i-1][j] != -1 {
					dp[i][j] = min(dp[i][j-coin]+1, dp[i-1][j])
				} else {
					// 否则，只能通过dp[i][j - coin]加上当前的coin得到amount总和
					dp[i][j] = dp[i][j-coin] + 1
				}
			}
		}
	}
	return dp[n][amount]
}

// 动态规划
// 自下而上
func coinChange(coins []int, amount int) int { // 凑成总金额所需的 最少的硬币个数
	// https://leetcode-cn.com/problems/coin-change/solution/322-ling-qian-dui-huan-by-leetcode-solution/
	dp := make([]int, amount+1) // dp数组，长度为amount+1
	// 初始化dp
	// 定义 dp(i) 为组成金额 i 所需最少的硬币数量
	dp[0] = 0                      // 当 i==0 时无法用硬币组成，为 0
	for i := 1; i < len(dp); i++ { // 从下标1开始
		dp[i] = amount + 1 // amount+1是因为后面用的是min?
	}

	for i := 1; i < amount; i++ { // 从1开始
		for j := 0; j < len(coins); j++ { // 枚举硬币的面值，这个for是关键
			if coins[j] <= i { // 硬币面值小于金额
				// 该硬币的面值 + dp[i-coins[j]]的和小于等于amount
				dp[i] = min(dp[i], dp[i-coins[j]]+1) // 要硬币数量最少，所以 dp[i] 为前面能转移过来的状态的最小值加上枚举的硬币数量 1
			}
		}
	}
	if dp[amount] > amount { // 重点
		return -1
	}
	return dp[amount]
}
