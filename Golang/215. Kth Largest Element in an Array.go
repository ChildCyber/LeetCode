package leetcode

import (
	"container/heap"
	"math/rand"
	"sort"
	"time"
)

// 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

/*
在快排的 partition 操作中，每次 partition 操作结束都会返回一个点，这个标定点的下标和最终排序之后有序数组中这个元素所在的下标是一致的。
利用这个特性，可以不断的划分数组区间，最终找到第 K 大的元素。
执行一次 partition 操作以后，如果这个元素的下标比 K 小，那么接着就在后边的区间继续执行 partition 操作;
如果这个元素的下标比 K 大，那么就在左边的区间继续执行 partition 操作;
如果相等就直接输出这个下标对应的数组元素即可。
*/
// 实现快排，取第K个元素
func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, left, right, index int) int {
	q := randomPartition(arr, left, right)
	if q == index {
		return arr[q]
	} else if q < index {
		return quickSelect(arr, q+1, right, index)
	}
	return quickSelect(arr, left, q-1, index)
}

func randomPartition(arr []int, left, right int) int {
	i := rand.Int()
	arr[i], arr[right] = arr[right], arr[i]
	return partition215(arr, left, right)
}

func partition215(arr []int, left, right int) int {
	x := arr[right]
	i := left - 1
	for j := left; j < right; j++ {
		if arr[j] <= x {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

// 标准库排序，取第K个元素
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 堆
// 大根堆
// 时间复杂度：O(n + k log n)
// 空间复杂度：O(n)
// 思路：把所有元素都放进去。弹出堆顶 k-1 次，第 k 次弹出的就是第 K 大
func findKthLargest2(nums []int, k int) int {
	heapSize := len(nums)
	// 创建大小为nums长度的大根堆
	buildMaxHeap(nums, heapSize)
	// 做 k−1 次删除操作
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	// 堆顶元素为答案
	return nums[0]
}

func buildMaxHeap(arr []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(arr, i, heapSize)
	}
}

func maxHeapify(arr []int, i, heapSize int) {
	left, right, largest := i*2+1, i*2+2, i
	if left < heapSize && arr[left] > arr[largest] {
		largest = left
	}
	if right < heapSize && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		maxHeapify(arr, largest, heapSize)
	}
}

type MaxHeap []int // 大根堆实现

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargestMaxHeap(nums []int, k int) int {
	// 使用大根堆，pop k-1次
	h := &MaxHeap{}
	*h = append(*h, nums...)
	heap.Init(h)

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return heap.Pop(h).(int)
}

// 小根堆
// 时间复杂度：O(n log k)
// 空间复杂度：O(k)
// 思路：维持大小为 k 的堆。每来一个新数，如果比堆顶大就替换

type MinHeap []int // 小根堆实现

// 实现sort.Interface接口
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// 实现heap.Interface接口
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func findKthLargestMinHeap(nums []int, k int) int {
	// 使用小根堆，保持堆大小为k
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return (*h)[0]
}
