package leetcode

// 最大子数组和
// https://leetcode-cn.com/problems/maximum-subarray/

// 暴力
func maxSubArrayForce(nums []int) int {
	n := len(nums)
	ans := nums[0] // 处理数组中所有元素都为负数的情况
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ { // [i:j]区间的子数组；j=i是为了避免子数组只包含一个元素的情况
			sum += nums[j]
			ans = max(sum, ans)
		}
	}
	return ans
}

// 动态规划
func maxSubArray(nums []int) int {
	// 状态转移方程： dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0) , dp[i] = nums[i] (dp[i-1] ≤ 0)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	dp, ans := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ { // 从1开始
		// dp的两种情况
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		ans = max(ans, dp[i]) // 最大子数组和
	}
	return ans
}

// 模拟
// kandane算法：以每个位置为结尾的最大子数组和，要么是当前元素本身，要么是当前元素加上以 *前一个位置为结尾* 的最大子数组和
func maxSubArrayAnalog(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// 想象成一个"贪心"的过程，逐个考虑数组中的元素：
	// 1. 如果之前的子数组和是负数，那么抛弃之前的子数组，从当前元素重新开始会更有利
	// 2. 如果之前的子数组和是正数，那么加上当前元素可能会使和更大
	// 3. 更新最大值
	globalMax, curMax := nums[0], 0
	for i := 0; i < len(nums); i++ {
		curMax += nums[i]
		if curMax > globalMax {
			globalMax = curMax
		}
		if curMax < 0 { // 贪心:局部最优
			curMax = 0
		}
	}
	return globalMax
}

// 分治
// 将数组分为左右两半
// 最大子数组和可能出现在：
// 左半部分
// 右半部分
// 跨越中间的部分
// 递归求解左半部分和右半部分的最大子数组和
// 计算跨越中间的最大子数组和（需要从中间向左右扩展）
// 取三者中的最大值
