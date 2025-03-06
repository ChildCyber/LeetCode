package leetcode

// 寻找旋转排序数组中的最小值
// https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/
// lc 33的简化版

// 直接遍历
// 时间复杂度：O(n)

// 二分查找
// 时间复杂度：O(log n)
func findMin(nums []int) int {
	// 思路：每次二分后，数组会被mid分成左右两部分，其中必有一部分是完全有序的
	// 先判断mid在哪一个有序区间，再决定收缩哪边的边界
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 判断最小值在哪一侧
		if nums[mid] < nums[right] { // mid及其右侧[mid...high]是有序递增的，最小值在左侧（包括mid）
			right = mid // 保留mid
		} else { // 最小值在右侧（不包括mid）
			left = mid + 1
		}
	}
	return nums[left]
}
