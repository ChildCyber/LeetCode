package leetcode

import (
	"fmt"
	"strconv"
)

// 汇总区间
// https://leetcode.cn/problems/summary-ranges/

// 双指针
func summaryRanges(nums []int) []string {
	// 使用双指针来标识每个连续区间的起点和终点：起点start，终点i
	if len(nums) == 0 {
		return []string{}
	}

	var ans []string
	start := 0

	for i := 0; i < len(nums); i++ {
		// 检查是否连续：当前元素+1等于下一个元素，且不是最后一个元素
		if i < len(nums)-1 && nums[i]+1 == nums[i+1] {
			continue
		}

		// 当前区间结束，构建区间字符串
		if start == i {
			ans = append(ans, fmt.Sprintf("%d", nums[start]))
		} else {
			ans = append(ans, fmt.Sprintf("%d->%d", nums[start], nums[i]))
		}

		// 更新下一个区间的起始位置
		start = i + 1
	}

	return ans
}

func summaryRanges1(nums []int) []string {
	var result []string
	n := len(nums)

	for i := 0; i < n; i++ {
		start := nums[i]

		// 寻找连续区间的终点
		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}

		// 构建区间字符串
		if start == nums[i] {
			result = append(result, strconv.Itoa(start))
		} else {
			result = append(result, fmt.Sprintf("%d->%d", start, nums[i]))
		}
	}

	return result
}
