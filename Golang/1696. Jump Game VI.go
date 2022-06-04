package leetcode

import "math"

// 跳跃游戏 VI
// https://leetcode.cn/problems/jump-game-vi/
// 动态规划
func maxResult1(nums []int, k int) int {
	dp := make([]int, len(nums))
	if k > len(nums) {
		k = len(nums)
	}

	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}

	for i := 1; i < len(nums); i++ {
		left, tmp := max(0, i-k), math.MinInt32 // i-k>=0防止越界
		for j := left; j < i; j++ {             // [i−k,i−1]区间中最大的值
			tmp = max(tmp, dp[j])
		}
		dp[i] = nums[i] + tmp
	}

	return dp[len(nums)-1]
}

// 单调队列
func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MinInt32
	}

	// 单调队列求滑动窗口最大值，求[i−k,i−1]滑动区间的最大值
	window := make([]int, k)
	// 单调队列里面存一个区间内最大值的下标
	// 1. 队列的队首永远都是最大值，队列从大到小降序排列。如果来了一个比队首更大的值的下标，需要将单调队列清空，只存这个新的最大值的下标
	// 2. 队列的⻓度为 k。从队尾插入新值，并把队头的最大值“挤”出队首

	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i] + dp[window[0]] // 队首
		// 存放最大值的下标
		for len(window) > 0 && dp[window[len(window)-1]] <= dp[i] {
			window = window[:len(window)-1] // 清除队尾
		}
		// 滑动窗口，保证窗口大小为k
		for len(window) > 0 && i-k >= window[0] {
			window = window[1:] // 不停清除队首，窗口不断向左滑动
		}
		window = append(window, i)
	}

	return dp[len(nums)-1]
}
