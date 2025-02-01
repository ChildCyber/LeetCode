package leetcode

// 寻找旋转排序数组中的最小值
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// 第33题的简化版

// 直接遍历
// 时间复杂度：O(n)

// 二分查找
// 时间复杂度：O(log n)
func findMin(nums []int) int {
	// 找出最小值（旋转点）
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] { // 最小值在左侧（包括mid），元素不重复无需考虑等于情况
			right = mid
		} else { // 最小值在右侧（不包括mid）
			left = mid + 1
		}
	}
	return nums[left]
}
