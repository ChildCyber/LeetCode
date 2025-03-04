package leetcode

import "container/heap"

// 雇佣 K 位工人的总代价
// https://leetcode.cn/problems/total-cost-to-hire-k-workers/

// 双指针+双堆
func totalCost(costs []int, k int, candidates int) int64 {
	n := len(costs)
	left, right := 0, n-1

	leftHeap := &IntHeap{}  // 维护前端的候选工人
	rightHeap := &IntHeap{} // 维护后端的候选工人
	heap.Init(leftHeap)
	heap.Init(rightHeap)

	// 初始化左右堆
	for left < candidates && left <= right {
		heap.Push(leftHeap, costs[left])
		left++
	}
	for right >= n-candidates && right >= left {
		heap.Push(rightHeap, costs[right])
		right--
	}

	var total int64
	for i := 0; i < k; i++ {
		// 若两堆都为空，说明候选人不够
		if leftHeap.Len() == 0 && rightHeap.Len() == 0 {
			break
		}

		// 比较堆顶最小值
		if rightHeap.Len() == 0 || (leftHeap.Len() > 0 && (*leftHeap)[0] <= (*rightHeap)[0]) {
			total += int64(heap.Pop(leftHeap).(int))
			// 从左侧补充下一个候选人
			if left <= right {
				heap.Push(leftHeap, costs[left])
				left++
			}
		} else {
			total += int64(heap.Pop(rightHeap).(int))
			// 从右侧补充下一个候选人
			if left <= right {
				heap.Push(rightHeap, costs[right])
				right--
			}
		}
	}
	return total
}
