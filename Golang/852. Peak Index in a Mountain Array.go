package leetcode

// 山脉数组的峰顶索引
// https://leetcode-cn.com/problems/peak-index-in-a-mountain-array/
// O(n)
func peakIndexInMountainArray(arr []int) int {
	for i := 1; ; i++ {
		if arr[i] > arr[i+1] {
			return i
		}
	}
}

// 二分，O(log n)
func peakIndexInMountainArray1(arr []int) int {
	low, high := 0, len(arr)-1
	for low < high {
		mid := low + (high-low)/2
		// 如果 mid 较大，则左侧存在峰值，high = mid
		if arr[mid] > arr[mid+1] {
			high = mid
		} else { // 如果 mid + 1 较大，则右侧存在峰值，low = mid + 1
			low = mid + 1
		}
	}
	return low
}

func peakIndexInMountainArray2(arr []int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)>>1
		if arr[mid] > arr[mid+1] && arr[mid] > arr[mid-1] { // mid处于峰值
			return mid
		}
		if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] { // mid处于下坡，需要向左走
			high = mid - 1
		}
		if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] { // mid处于上坡，需要往右走
			low = mid + 1
		}
	}
	return 0
}
