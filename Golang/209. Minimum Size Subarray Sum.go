package leetcode

import "math"

// 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 连续子数组

// 暴力法
func minSubArrayLen(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum >= s {
				ans = min(ans, j-i+1)
				break
			}
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 滑动窗口
func minSubArrayLen1(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	start, end := 0, 0
	sum := 0
	for end < n { // 小于，不可越界
		sum += nums[end]
		for sum >= s { // 大于等于，start右移
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
