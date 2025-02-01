package leetcode

import "sort"

// 在排序数组中查找元素的第一个和最后一个位置
// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
// 二分
func searchRange(nums []int, target int) []int {
	// 第一次二分查找：专门用来找左边界（第一个等于目标值的位置）
	// 第二次二分查找：专门用来找右边界（最后一个等于目标值的位置）
	return []int{searchFirstEqualElement(nums, target), searchLastEqualElement(nums, target)}
}

// 二分查找第一个与 target 相等的元素
// 时间复杂度 O(log n)
func searchFirstEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else { // 上面一样，只有找到target后有变化
			if (mid == 0) || (nums[mid-1] != target) { // 检查是否越界，以及该位置的值是否等于target
				return mid // 找到第一个与 target 相等的元素
			}
			high = mid - 1 // 找到target，但不立即返回，继续在左侧搜索
		}
	}
	return -1
}

// 二分查找最后一个与 target 相等的元素
// 时间复杂度 O(log n)
func searchLastEqualElement(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + ((high - low) >> 1)

		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if (mid == len(nums)-1) || (nums[mid+1] != target) { // 检查是否越界，以及该位置的值是否等于target
				return mid // 找到最后一个与 target 相等的元素
			}
			low = mid + 1 // 找到target，但不立即返回，继续在右侧搜索
		}
	}
	return -1
}

// 利用标准库进行二分查找
func searchRange1(nums []int, target int) []int {
	// sort.SearchInts函数本身就是基于二分查找实现的，函数返回的是第一个大于等于target的元素索引
	leftmost := sort.SearchInts(nums, target) // 第一个
	if leftmost == len(nums) || nums[leftmost] != target {
		return []int{-1, -1}
	}

	rightmost := sort.SearchInts(nums, target+1) - 1 // 最后一个
	return []int{leftmost, rightmost}
}
