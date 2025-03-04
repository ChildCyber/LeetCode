package leetcode

import (
	"math"
	"math/rand"
)

// 寻找峰值
// https://leetcode-cn.com/problems/find-peak-element/
// 数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可
// 时间复杂度为 O(log n)

// 迭代爬坡
// 寻找最大值，数组 nums 中最大值两侧的元素一定严格小于最大值本身。因此，最大值所在的位置就是一个可行的峰值位置
// 时间复杂度：O(n)
func findPeakElementN(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			// 当前元素比右边大，说明找到了一个峰值
			return i
		}
		// 否则继续向右爬坡
	}
	// 如果一直上坡到最后，最后一个元素就是峰值
	return len(nums) - 1
}

// 迭代爬坡+随机化
// 在 [0,n) 的范围内随机一个初始位置 i，随后根据 nums[i−1],nums[i],nums[i+1] 三者的关系决定向哪个方向走：
// 如果 nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，可以直接返回 i；
// 如果 nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，需要往右走，即 i←i+1；
// 如果 nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，需要往左走，即 i←i−1；
// 如果 nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，可以朝任意方向走。
func findPeakElementPa(nums []int) int {
	n := len(nums)
	// 随机选择一个起始位置，而不是固定从开头开始
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

// 二分
// 从一个位置开始，不断地向高处走，那么最终一定可以到达一个峰值位置
// 时间复杂度：O(log n)
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right { // 确保在只剩一个元素时退出循环
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			// 处于上坡，峰值在右侧（包括mid+1）
			left = mid + 1
		} else {
			// 处于下坡或峰顶，峰值在左侧（包括mid本身）
			right = mid
		}
	}
	// 循环结束时 left == right，指向的就是一个峰值
	return left
}
