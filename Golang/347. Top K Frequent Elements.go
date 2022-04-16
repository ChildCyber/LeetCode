package leetcode

import "container/heap"

// 前 K 个高频元素
// https://leetcode-cn.com/problems/top-k-frequent-elements/
type Item struct {
	key   int
	count int
}

type PriorityQueue []*Item

// 实现heap.Interface接口
// sort.Interface
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count > pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// heap.Push
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

// heap.Pop
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	*pq = (*pq)[:n-1]
	return item
}

// 哈希表 + 堆
// 时间复杂度：O(N log k)，其中 N 为数组的长度，k 为堆的大小
// 空间复杂度：O(N)
func topKFrequent(nums []int, k int) []int {
	// 遍历整个数组，使用哈希表记录每个数字出现的次数，构造出现次数数组
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	// 建立一个小根堆，遍历「出现次数数组」
	q := PriorityQueue{}
	for key, count := range m {
		heap.Push(&q, &Item{
			key:   key,
			count: count,
		})
	}

	// 取出堆中前 k 大的值
	var result []int
	for len(result) < k {
		item := heap.Pop(&q).(*Item)
		result = append(result, item.key)
	}
	return result
}
