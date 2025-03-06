package leetcode

import "math"

// 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
// 连续子数组

// 暴力
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func minSubArrayLenBrute(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	ans := math.MaxInt32
	// 枚举所有子数组：对每个起点i，遍历终点j，计算sum，看是否满足≥target
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
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func minSubArrayLen(target int, nums []int) int {
	// 思路：每一轮迭代，将 nums[end]加到 sum，
	// 如果 sum≥s，则更新子数组的最小长度（此时子数组的长度是end−start+1），
	// 然后将 nums[start] 从 sum 中减去并将 start 右移，直到 sum<s，在此过程中同样更新子数组的最小长度。
	// 在每一轮迭代的最后，将 end 右移。
	n := len(nums)
	left, sum := 0, 0
	minLen := n + 1

	for right := 0; right < n; right++ {
		sum += nums[right] // 扩大窗口

		// 当窗口内的和 >= target 时，尝试收缩
		for sum >= target {
			currLen := right - left + 1
			if currLen < minLen {
				minLen = currLen
			}
			sum -= nums[left] // 缩小窗口
			left++
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}
