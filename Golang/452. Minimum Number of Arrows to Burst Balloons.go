package leetcode

import "sort"

// 用最少数量的箭引爆气球
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/

// 排序+遍历
func findMinArrowShots(points [][]int) int {
	n := len(points)
	if n == 0 {
		return 0
	}

	// 1. 按右端点排序（升序）
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	// 2. 初始化
	arrows := 1         // 至少需要一箭
	pos := points[0][1] // 当前箭射在第一个区间的右端点

	// 3. 遍历区间
	for _, p := range points {
		if p[0] > pos {
			// 当前区间与上一个不重叠 -> 需要新箭
			arrows++
			pos = p[1] // 更新射箭位置
		}
		// 否则：说明 p[0] <= pos，有重叠，一箭足矣
	}

	return arrows
}
