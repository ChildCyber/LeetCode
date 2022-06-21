package leetcode

// 最大子数组和
// https://leetcode-cn.com/problems/maximum-subarray/

// 暴力
func maxSubArrayForce(nums []int) int {
	n := len(nums)
	ans := nums[0]
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ { // [i:j]区间的子数组
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
// 不用建立dp数组，直接用sum变量来记录dp[i−1]的答案
func maxSubArrayAnalog(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	// 1. 如果之前的和<0，则重新计算
	// 2. 如果之前的和>=0，则计入sum结果中
	// 3. 更新最大值
	ans, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum > ans {
			ans = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return ans
}

// 分治
