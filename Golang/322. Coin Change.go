package leetcode

import (
	"math"
	"sort"
)

// 零钱兑换
// https://leetcode-cn.com/problems/coin-change/
// 返回凑成总金额所需的最少的硬币个数

// 暴力递归（DFS）
func coinChangeRec(coins []int, amount int) int {
	// 边界情况：如果没有硬币，直接返回-1
	if len(coins) == 0 {
		return -1
	}

	// 初始化答案为最大整数值，用于后续比较找到最小值
	ans := math.MaxInt32

	// 定义递归函数，使用闭包可以访问外部变量ans
	var findWay func([]int, int, int)
	findWay = func(coins []int, amount int, count int) {
		// 基准情况1：剩余金额小于0，说明这条路径不可行，直接返回
		if amount < 0 {
			return
		}
		// 基准情况2：剩余金额等于0，找到一种可行的硬币组合
		if amount == 0 {
			// 更新最小硬币数量
			ans = min(ans, count)
		}
		// 递归情况：尝试使用每一种硬币
		for i := 0; i < len(coins); i++ {
			// 选择当前硬币，减少相应金额，增加硬币计数
			findWay(coins, amount-coins[i], count+1)
		}
	}

	// 从金额为amount，硬币数为0开始递归搜索
	findWay(coins, amount, 0)
	// 如果ans没有被更新过，说明没有找到解
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 深度优先搜索+剪枝（DFS+Pruning）
func coinChangeRecPruning(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}

	// 对硬币进行从大到小排序，优先使用大面值硬币，有助于更快剪枝
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	ans := math.MaxInt32

	var findWay func([]int, int, int)
	findWay = func(coins []int, amount int, count int) {
		// 剪枝1：如果当前硬币数已经大于等于已知最优解，直接返回
		if count >= ans {
			return
		}

		// 基准情况1：剩余金额小于0，路径不可行
		if amount < 0 {
			return
		}

		// 基准情况2：找到一种可行解
		if amount == 0 {
			ans = count
			return
		}

		// 递归尝试每种硬币
		for i := 0; i < len(coins); i++ {
			// 剪枝2：如果硬币面值大于剩余金额，跳过
			if coins[i] > amount {
				continue
			}

			// 剪枝3：预估最小硬币数 = 当前硬币数 + 剩余金额/当前硬币面值（向上取整）
			// 如果这个预估值已经大于等于已知最优解，跳过
			estimatedMin := count + (amount+coins[i]-1)/coins[i]
			if estimatedMin >= ans {
				continue
			}

			findWay(coins, amount-coins[i], count+1)
		}
	}

	findWay(coins, amount, 0)

	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 广度优先搜索（BFS）
func coinChangeBFS(coins []int, amount int) int {
	// 特殊情况处理：目标金额为0，不需要任何硬币
	if amount == 0 {
		return 0
	}

	// 创建队列，用于BFS遍历，存储当前金额
	queue := []int{amount}

	// 创建访问标记数组，避免重复计算相同金额
	// visited[i]表示金额i是否已被访问过
	visited := make([]bool, amount+1)
	visited[amount] = true // 标记起始节点已访问

	// 层级计数器，表示当前使用的硬币数
	level := 0

	// BFS主循环
	for len(queue) > 0 {
		// 增加硬币数（BFS层级）
		level++

		// 记录当前层级的节点数量
		levelSize := len(queue)

		// 遍历当前层级的所有节点
		for i := 0; i < levelSize; i++ {
			// 从队列头部取出当前金额
			current := queue[0]
			queue = queue[1:] // 移除已处理的节点

			// 尝试使用每一种硬币
			for _, coin := range coins {
				// 计算使用当前硬币后的新金额
				nextAmount := current - coin

				// 如果新金额为0，找到最优解
				if nextAmount == 0 {
					return level
				}

				// 如果新金额有效且未被访问过
				if nextAmount > 0 && !visited[nextAmount] {
					// 标记为已访问，避免重复处理
					visited[nextAmount] = true
					// 将新金额加入队列，等待下一层处理
					queue = append(queue, nextAmount)
				}
			}
		}
	}

	// 如果BFS完成仍未找到解，返回-1
	return -1
}

// 动态规划-自顶向下（记忆化搜索）
func coinChangeMemo(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
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

// 动态规划：二维DP
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
			// 如果硬币面值大于amount，或者是当前amount不能由上一个 j-coin 得到，则从上一个状态转移
			if coin > j || dp[i][j-coin] == -1 {
				dp[i][j] = dp[i-1][j]
			} else {
				// 如果dp[i-1][j] != -1说明上一个状态可以凑成amount
				if dp[i-1][j] != -1 {
					dp[i][j] = min(dp[i][j-coin]+1, dp[i-1][j])
				} else {
					// 否则，只能通过dp[i][j-coin]加上当前的coin得到amount总和
					dp[i][j] = dp[i][j-coin] + 1
				}
			}
		}
	}
	return dp[n][amount]
}

// 动态规划
// 自下而上，一维DP
func coinChange(coins []int, amount int) int {
	// 定义状态：dp[i] 为组成金额 i 所需最少的硬币数量
	// 状态转移：对于当前要凑的金额 i，遍历每一种硬币面值 coin
	//   如果 coin 小于等于 i，说明这枚硬币可能被使用
	//   那么凑出金额 i 的一个可能方案就是：先凑出金额 i - coin，然后再加上这一枚硬币
	//   所以 dp[i] 的值就应该看看所有可能的 coin 中，哪个能使得 dp[i - coin] + 1 的值最小
	dp := make([]int, amount+1) // dp数组，长度为amount+1

	// 初始化dp
	dp[0] = 0                      // 金额为0需要0个硬币
	for i := 1; i <= amount; i++ { // 从下标1开始
		dp[i] = amount + 1 // 用于最后判断是否能找零
	}

	for i := 1; i <= amount; i++ { // 从1开始，当前的金额amount
		for j := 0; j < len(coins); j++ { // 当前硬币的面值
			if coins[j] <= i { // 可以使用当前的硬币进行找零
				dp[i] = min(dp[i], dp[i-coins[j]]+1) // 要硬币的数量最少，所以 dp[i] 为前面能转移过来的状态的最小值加上枚举的硬币数量 1
			}
		}
	}

	// 如果dp[amount]大于amount，说明无法凑出
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
