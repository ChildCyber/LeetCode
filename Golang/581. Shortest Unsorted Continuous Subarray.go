package leetcode

import (
	"math"
	"sort"
)

// 最短无序连续子数组
// https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/

// 排序
// 将给定的数组 nums 表示为三段子数组拼接的形式，分别记作 numsA​，numsB​，numsC
// 要求找到最短的 numsB，即最短逆序区间
// 将原数组 nums 排序与原数组进行比较，取最长的相同的前缀为 numsA​，取最长的相同的后缀为 numsC​，这样就可以取到最短的 numsB
func findUnsortedSubarraySort(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}

	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted) // 排序

	left, right := 0, len(nums)-1
	// 取最长的相同的前缀为 numsA​
	for nums[left] == numsSorted[left] {
		left++
	}
	// 取最长的相同的后缀为 numsC
	for nums[right] == numsSorted[right] {
		right--
	}

	return right - left + 1
}

// 一次遍历
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	// 假设 numsB 在 nums 中对应区间为 [left,right]
	left, right := -1, -1
	// 从大到小枚举 i，用一个变量 minn 记录，left 左侧即为 numsA 能取得的最大范围
	// 从大到小枚举 i，用一个变量 maxn 记录，right 右侧即为 numsC 能取得的最大范围
	minn, maxn := math.MaxInt64, math.MinInt64

	for i, num := range nums {
		if maxn > num {
			right = i
		} else {
			maxn = num
		}

		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}

	if right == -1 {
		return 0
	}

	return right - left + 1
}
