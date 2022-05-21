package leetcode

// 乘积最大子数组
// https://leetcode-cn.com/problems/maximum-product-subarray/
// 动态规划
func maxProduct(nums []int) int {
	// 维护两个变量，当前的最大值，以及最小值，最小值可能为负数
	minimum, maximum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 { // 乘以一个负数，当前的最大值就变成最小值，而最小值则变成最大值了
			maximum, minimum = minimum, maximum
		}
		maximum = max(nums[i], maximum*nums[i])
		minimum = min(nums[i], minimum*nums[i])
		res = max(res, maximum)
	}
	return res
}
