package leetcode

import "sort"

// 插入区间
// https://leetcode.cn/problems/insert-interval/

// 排序
func insertSort(intervals [][]int, newInterval []int) [][]int {
	// 追加newInterval，简化成lc 56
	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	for _, cur := range intervals[1:] {
		last := ans[len(ans)-1]
		if cur[0] <= last[1] {
			if cur[1] > last[1] {
				ans[len(ans)-1][1] = cur[1]
			}
		} else {
			ans = append(ans, cur)
		}
	}
	return ans
}

// 排序
func insert(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)
	i, n := 0, len(intervals)

	// 阶段1：添加所有结束点小于新区间起始点的区间
	for i < n && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}

	// 阶段2：合并所有与新区间有重叠的区间
	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	ans = append(ans, newInterval)

	// 阶段3：添加剩余的区间
	for i < n {
		ans = append(ans, intervals[i])
		i++
	}

	return ans
}
