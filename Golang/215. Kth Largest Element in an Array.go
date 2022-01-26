package leetcode

import (
	"math/rand"
	"sort"
	"time"
)

// 数组中的第K个最大元素
// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
/*
在快排的 partition 操作中，每次 partition 操作结束都会返回一个点，这个标定点的下标和最终排序之
后有序数组中这个元素所在的下标是一致的。利用这个特性，我们可以不断的划分数组区间，最终找到
第 K 大的元素。执行一次 partition 操作以后，如果这个元素的下标比 K 小，那么接着就在后边的区间
继续执行 partition 操作;如果这个元素的下标比 K 大，那么就在左边的区间继续执行 partition 操作;
如果相等就直接输出这个下标对应的数组元素即可。
*/
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
	return partition(arr, left, right)
}

func partition(arr []int, left, right int) int {
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

// 直接排序，取K
func findKthLargest1(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 建立一个大根堆，做 k−1 次删除操作后堆顶元素就是答案
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
