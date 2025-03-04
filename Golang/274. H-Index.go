package leetcode

import "sort"

// H 指数
// https://leetcode.cn/problems/h-index/

// 排序+扫描
func hIndex(citations []int) int {
	n := len(citations)
	if n == 0 {
		return 0
	}

	// 1. 排序
	sort.Ints(citations)

	// 2. 从后往前扫描寻找H指数
	h := 0
	for i := n - 1; i >= 0; i-- {
		// 当前有 n-i 篇论文引用次数 >= citations[i]
		// 需要 citations[i] >= n-i
		if citations[i] >= n-i {
			h = n - i
		} else {
			break
		}
	}

	return h
}

// 计数排序
// 对于引用次数可能很大的情况，可以使用计数排序优化到 O(n) 时间复杂度
func hIndex2(citations []int) int {
	sort.Ints(citations)
	n := len(citations)

	for i := 0; i < n; i++ {
		// 有 n-i 篇论文引用次数 >= citations[i]
		if citations[i] >= n-i {
			return n - i
		}
	}

	return 0
}
