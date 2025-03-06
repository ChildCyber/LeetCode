package leetcode

import "sort"

// 合并区间
// https://leetcode-cn.com/problems/merge-intervals/
// 排序
func merge56(intervals [][]int) [][]int {
	// 从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{}
	for i := 0; i < len(intervals); i++ {
		if len(ans) == 0 || ans[len(ans)-1][1] < intervals[i][0] { // 如果列表为空，或者当前区间与上一区间不重合，直接添加
			ans = append(ans, intervals[i])
		} else { // 与上一区间进行合并，更新上一区间的右端点
			ans[len(ans)-1][1] = max(ans[len(ans)-1][1], intervals[i][1])
		}
	}
	return ans
}

func merge56_1(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ans := [][]int{intervals[0]}
	for _, cur := range intervals[1:] { // 从索引1开始
		last := ans[len(ans)-1]
		if cur[0] <= last[1] { // 判断重叠
			if cur[1] > last[1] { // 合并区间：处理当前区间部分重叠且延伸更远
				ans[len(ans)-1][1] = cur[1]
			}
		} else {
			ans = append(ans, cur)
		}
	}
	return ans
}
