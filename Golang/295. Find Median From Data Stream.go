package leetcode

import (
	"container/heap"
)

// 数据流的中位数
// https://leetcode.cn/problems/find-median-from-data-stream/

// MedianFinder 优先队列
type MedianFinder struct {
	maxHeap *MaxHeap // 大根堆，存储较小的一半
	minHeap *MinHeap // 小根堆，存储较大的一半
}

func Constructor295() MedianFinder {
	maxHeap := &MaxHeap{}
	minHeap := &MinHeap{}
	heap.Init(maxHeap)
	heap.Init(minHeap)
	return MedianFinder{maxHeap: maxHeap, minHeap: minHeap}
}

func (mf *MedianFinder) AddNum(num int) {
	// 平衡策略：
	//   始终保持 maxHeap 的大小等于或比 minHeap 大 1
	//   确保 maxHeap 的最大值 <= minHeap 的最小值
	// 先加到最大堆
	heap.Push(mf.maxHeap, num)
	// 确保最大堆的最大值 <= 最小堆的最小值
	heap.Push(mf.minHeap, heap.Pop(mf.maxHeap))
	// 平衡两边大小
	if mf.minHeap.Len() > mf.maxHeap.Len() {
		heap.Push(mf.maxHeap, heap.Pop(mf.minHeap))
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	// 奇数个元素：中位数在大根堆堆顶
	if mf.maxHeap.Len() > mf.minHeap.Len() {
		return float64((*mf.maxHeap)[0])
	}
	// 偶数个元素：中位数是两个堆顶的平均值
	return (float64((*mf.maxHeap)[0]) + float64((*mf.minHeap)[0])) / 2.0
}
