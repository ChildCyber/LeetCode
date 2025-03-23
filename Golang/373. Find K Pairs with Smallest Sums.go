package leetcode

import "container/heap"

// 查找和最小的 K 对数字
// https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/

// 小根堆
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 思路：
	// 初始将每个 nums1[i] 与 nums2[0] 组成的数对加入堆
	// 每次从堆中取出最小数对 (i, j)
	// 然后将 (i, j+1) 加入堆（如果 j+1 有效）
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return [][]int{}
	}

	ans := make([][]int, 0, k)
	h := &MinHeap373{}
	heap.Init(h)

	// 初始化：将每个nums1[i]与nums2[0]的组合加入堆
	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, []int{nums1[i] + nums2[0], i, 0}) // 堆里存储的是(和,nums1的索引,nums2的索引)
	}

	// 取出前k小的数对
	for h.Len() > 0 && len(ans) < k {
		item := heap.Pop(h).([]int)
		i, j := item[1], item[2]
		ans = append(ans, []int{nums1[i], nums2[j]})

		// 将下一个候选加入堆
		if j+1 < len(nums2) {
			heap.Push(h, []int{nums1[i] + nums2[j+1], i, j + 1})
		}
	}

	return ans
}

type MinHeap373 [][]int

func (h MinHeap373) Len() int           { return len(h) }
func (h MinHeap373) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap373) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap373) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap373) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
