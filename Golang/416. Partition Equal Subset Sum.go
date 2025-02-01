package leetcode

// 分割等和子集
// https://leetcode.cn/problems/partition-equal-subset-sum/
// 动态规划
func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	// 计算总和
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// 如果总和是奇数，直接返回false
	if sum%2 != 0 {
		return false
	}

	// 目标值
	target := sum / 2

	// 状态定义：定义dp[i][j]为布尔值，表示考虑前i个数字（nums[0]到nums[i-1]），能否选出一些数字，使它们的和正好等于j
	// 状态转移：
	//   情况1：不选第i个数字
	//     如果前i-1个数字能凑出j，那么前i个数字也能凑出j，dp[i][j] = dp[i-1][j]
	//   情况2：选第i个数字
	//     前提：j >= nums[i-1]（目标和要大于等于当前数字）
	//     如果前i-1个数字能凑出 j-nums[i-1]，那么加上当前数字就能凑出j，dp[i][j] = dp[i-1][j-nums[i-1]]
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, target+1)
	}
	// 初始化：不选任何数字时，和为0是可以达到的
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			currentNum := nums[i-1]
			if j < currentNum {
				// 当前数字太大，不能选
				dp[i][j] = dp[i-1][j]
			} else {
				// 可以选择当前数字或不选
				dp[i][j] = dp[i-1][j] || dp[i-1][j-currentNum]
			}
		}
	}
	return dp[n][target]
}

// 动态规划一维数组实现（空间优化）
func canPartitionOptimized(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return true
	}

	// 计算总和
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// 如果总和是奇数，直接返回false
	if sum%2 != 0 {
		return false
	}

	target := sum / 2

	// 创建一维dp数组
	dp := make([]bool, target+1)
	dp[0] = true // 和为0总是可以达到

	// 遍历每个数字
	for i := 0; i < n; i++ {
		currentNum := nums[i]

		// 从后往前遍历，避免重复使用同一个数字
		for j := target; j >= currentNum; j-- {
			// 如果j-currentNum可以达到，那么j也可以达到
			if dp[j-currentNum] {
				dp[j] = true
			}
		}

		// 如果已经找到目标值，可以提前结束
		if dp[target] {
			return true
		}
	}

	return dp[target]
}
