package leetcode

import "sort"

// 无重叠区间
// https://leetcode.cn/problems/non-overlapping-intervals/

// 排序+双指针
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// 按区间结束时间排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 初始化计数器和当前区间的结束时间
	count := 1 // 记录最多能保留的不重叠区间数（第一个区间肯定保留）
	end := intervals[0][1]

	// 遍历所有区间
	for i := 1; i < len(intervals); i++ {
		// 如果当前区间的开始时间 >= 前一个保留区间的结束时间
		// 说明不重叠，可以保留
		if intervals[i][0] >= end {
			count++
			end = intervals[i][1] // 更新结束时间
		}
		// 否则就是重叠，跳过（相当于移除）
	}

	// 需要移除的区间数 = 总区间数 - 保留的区间数
	return len(intervals) - count
}
