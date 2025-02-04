package leetcode

import (
	"container/heap"
	"sort"
)

// 前 K 个高频元素
// https://leetcode-cn.com/problems/top-k-frequent-elements/

// 哈希表+排序
// 时间复杂度：
//   频率统计：O(n)
//   排序：O(m log m)，其中m是不同元素的数量
// 总体：O(n + m log m)
// 空间复杂度：O(m)，其中m是不同元素的数量
func topKFrequentSort(nums []int, k int) []int {
	// 第一步：使用map统计每个元素的频率
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	// 第二步：将map转换为键值对切片，方便排序
	pairs := make([][2]int, 0, len(freqMap))
	for num, freq := range freqMap {
		pairs = append(pairs, [2]int{num, freq})
	}

	// 第三步：按频率降序排序
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] > pairs[j][1] // 按频率降序排列
	})

	// 第四步：取前k个元素
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = pairs[i][0]
	}

	return result
}

// 哈希表+小根堆
// 时间复杂度：O(N log k)，其中 N 为数组的长度，k 为堆的大小
// 空间复杂度：O(N)
func topKFrequent(nums []int, k int) []int {
	// 遍历整个数组，使用哈希表记录每个数字出现的次数，构造出现次数数组
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}

	// 创建最小堆并初始化
	q := &PriorityQueue{}
	heap.Init(q)
	for key, count := range m {
		heap.Push(q, &Item{
			num:  key,
			freq: count,
		})
		if q.Len() > k {
			heap.Pop(q)
		}
	}

	// 取出堆中前 k 大的值
	var result []int
	for i := k - 1; i >= 0; i-- {
		item := heap.Pop(q).(*Item)
		result = append(result, item.num)
	}
	return result
}

type Item struct {
	num  int
	freq int
}

type PriorityQueue []*Item

// 实现sort.Interface接口
func (h PriorityQueue) Len() int {
	return len(h)
}

func (h PriorityQueue) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 实现heap.Interface接口
func (h *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*h = append(*h, item)
}

func (h *PriorityQueue) Pop() interface{} {
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}
