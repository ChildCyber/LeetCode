package leetcode

// 零钱兑换
// https://leetcode-cn.com/problems/coin-change/
// 动态规划
func coinChange(coins []int, amount int) int { // 凑成总金额所需的 最少的硬币个数
	// 采用自下而上的方式进行思考
	// 三角形
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
