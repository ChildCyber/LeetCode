package leetcode

import "container/heap"

// 无限集中的最小数字
// https://leetcode.cn/problems/smallest-number-in-infinite-set/

// 堆+哈希表
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SmallestInfiniteSet struct {
	nextNum int          // 下一个未取的自然数
	minHeap *IntHeap     // 存储加回的数字
	inHeap  map[int]bool // 去重
}

func Constructor2336() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)
	return SmallestInfiniteSet{
		nextNum: 1,
		minHeap: h,
		inHeap:  make(map[int]bool),
	}
}

func (s *SmallestInfiniteSet) PopSmallest() int {
	if s.minHeap.Len() > 0 {
		val := heap.Pop(s.minHeap).(int)
		delete(s.inHeap, val)
		return val
	}
	val := s.nextNum
	s.nextNum++
	return val
}

func (s *SmallestInfiniteSet) AddBack(num int) {
	if num < s.nextNum && !s.inHeap[num] {
		heap.Push(s.minHeap, num)
		s.inHeap[num] = true
	}
}
