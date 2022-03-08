package leetcode

// 寻找旋转排序数组中的最小值
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// 遍历，O(n)

// 二分，O(log n)
func findMin(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)/2
		if nums[mid] < nums[high] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return nums[low]
}
