package leetcode

import (
	"math"
	"math/rand"
)

// 寻找峰值
// https://leetcode-cn.com/problems/find-peak-element/
// 数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可
// 时间复杂度为 O(log n)

// 寻找最大值，数组 nums 中最大值两侧的元素一定严格小于最大值本身。因此，最大值所在的位置就是一个可行的峰值位置
// 时间复杂度 O(n)
func findPeakElementN(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}

// 迭代爬坡
// 在 [0,n) 的范围内随机一个初始位置 i，随后根据 nums[i−1],nums[i],nums[i+1] 三者的关系决定向哪个方向走：
// 如果 nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，可以直接返回 i；
// 如果 nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，需要往右走，即 i←i+1；
// 如果 nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，需要往左走，即 i←i−1；
// 如果 nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，可以朝任意方向走。
func findPeakElementPa(nums []int) int {
	n := len(nums)
	idx := rand.Intn(n)

	// 辅助函数，输入下标 i，返回 nums[i] 的值
	// 方便处理 nums[-1] 以及 nums[n] 的边界情况
	get := func(i int) int {
		if i == -1 || i == n {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx) < get(idx+1) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}

// 二分，迭代爬坡
// 从一个位置开始，不断地向高处走，那么最终一定可以到达一个峰值位置
// 时间复杂度 O(log n)
func findPeakElement(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high { // low < high
		mid := low + (high-low)/2
		// 如果 mid 较大，则左侧存在峰值，high = mid
		if nums[mid] > nums[mid+1] {
			high = mid
		} else { // 如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
			low = mid + 1
		}
	}
	return low
}
