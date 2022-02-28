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
	// 思路：每一轮迭代，将 nums[end]加到 sum，
	// 如果 sum≥s，则更新子数组的最小长度（此时子数组的长度是end−start+1），
	// 然后将 nums[start] 从 sum 中减去并将 start 右移，直到 sum<s，在此过程中同样更新子数组的最小长度。
	// 在每一轮迭代的最后，将 end 右移。
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	start, end := 0, 0 // 滑动窗口的开始位置和结束位置
	sum := 0           // 从nums[start]到nums[end]的子数组中的元素和
	for end < n {      // 小于，不可越界
		sum += nums[end]
		for sum >= s { // 大于等于，start右移
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++ // end右移
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
