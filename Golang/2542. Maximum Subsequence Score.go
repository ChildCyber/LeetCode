package leetcode

import (
	"container/heap"
	"sort"
)

// 最大子序列的分数
// https://leetcode.cn/problems/maximum-subsequence-score/

// 排序+小根堆
func maxScore(nums1 []int, nums2 []int, k int) int64 {
	n := len(nums1)

	// 创建pair数组并按照nums2降序排序
	pairs := make([][2]int, n)
	for i := 0; i < n; i++ {
		pairs[i] = [2]int{nums1[i], nums2[i]}
	}

	// 按nums2降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1]
	})

	// 初始化最小堆和当前和
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	var currentSum int64
	// 先添加前k-1个元素
	for i := 0; i < k-1; i++ {
		heap.Push(minHeap, pairs[i][0])
		currentSum += int64(pairs[i][0])
	}

	var ans int64
	// 从第k-1个元素开始遍历
	for i := k - 1; i < n; i++ {
		// 添加当前元素
		heap.Push(minHeap, pairs[i][0])
		currentSum += int64(pairs[i][0])

		// 如果堆大小等于k，计算当前分数
		if minHeap.Len() == k {
			currentScore := currentSum * int64(pairs[i][1]) // pairs[i][1]为当前最小
			if currentScore > ans {
				ans = currentScore
			}
		}

		// 移除最小的元素以保持堆大小为k-1（为下一次迭代准备）
		smallest := heap.Pop(minHeap).(int)
		currentSum -= int64(smallest)
	}

	return ans
}
