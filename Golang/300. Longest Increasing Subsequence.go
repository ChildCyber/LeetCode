package leetcode

// 最长递增子序列
// https://leetcode-cn.com/problems/longest-increasing-subsequence/
// 动态规划
func lengthOfLIS(nums []int) int {
	// dp[i] 代表 [0,i] 范围内,选择数字 nums[i] 为结尾可以获得的最⻓上升子序列的⻓度。
	// 状态转移方程为 dp[i] = max(dp[j]) + 1 ,其中 j < i && nums[j] > nums[i] ,取所有满足条件的最大值。
	dp, res := make([]int, len(nums)+1), 0 // 多了个dp[0]，长度需要加1
	dp[0] = 0                              // 不选数字，最⻓上升子序列的⻓度为0
	for i := 1; i <= len(nums); i++ {      // 需要两个for
		for j := 1; j < i; j++ { // j<i
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1 // i == 1, dp[1]=1
		res = max(res, dp[i])
	}
	return res
}

func lengthOfLIS1(nums []int) int {
	ans, length := 0, len(nums)
	if length == 0 {
		return ans
	}
	dp := make([]int, length) // dp长度和原数组相同
	for i := 0; i < length; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ { // j < i
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			ans = max(dp[i], ans)
		}
	}
	return ans
}
