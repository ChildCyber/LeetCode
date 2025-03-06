package leetcode

// 搜索插入位置
// https://leetcode.cn/problems/search-insert-position/
// 必须使用时间复杂度为 O(log n) 的算法

// 二分搜索
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1 // 经典二分，闭区间
	for left <= right {
		mid := left + (right-left)/2
		// 要找到第一个不小于target的位置
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 开区间
// 思路同标准库的sort.Search
// idx := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
func searchInsert1(nums []int, target int) int {
	left, right := 0, len(nums) // 答案二分，开区间
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else { // nums[mid]>=target，说明插入点在mid或之前
			right = mid
		}
	}
	// 循环终止时left==right，即为答案（第一个>=target的索引，或n）
	return left
}
