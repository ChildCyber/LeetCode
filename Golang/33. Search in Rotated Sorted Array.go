package leetcode

// 搜索旋转排序数组
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array
// 算法时间复杂度必须是 O(log n)

// 二分查找
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环。
func search33(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		}

		// 在常规二分查找的时候查看当前 mid 为分割位置分割出来的两个部分 [l, mid] 和 [mid + 1, r] 哪个部分是有序的
		// 并根据有序的那个部分确定该如何改变二分查找的上下界
		if nums[low] <= nums[mid] { // [low, mid - 1]为升序区间
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1 // 在 [low, mid - 1] 中寻找
			} else {
				low = mid + 1 // 在 [mid + 1, high] 中寻找
			}
		} else { // [mid, high]为升序区间
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1 // 在 [mid + 1, high] 中寻找
			} else {
				high = mid - 1 // 在 [l, mid - 1] 中寻找
			}
		}
	}
	return -1
}

// 先找旋转点，再二分查找
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// 第一步：找到旋转点（最小值的位置）
	pivot := findPivot(nums)

	// 第二步：根据目标值决定在哪一部分进行二分查找
	if pivot == 0 {
		// 数组没有旋转或旋转了一整圈
		return binarySearch(nums, 0, len(nums)-1, target)
	}

	if target >= nums[0] {
		// 目标值在旋转点的左侧（较大的部分）
		return binarySearch(nums, 0, pivot-1, target)
	} else {
		// 目标值在旋转点的右侧（较小的部分）
		return binarySearch(nums, pivot, len(nums)-1, target)
	}
}

// 找到旋转点（数组中最小值的位置）|二分查找
func findPivot(nums []int) int {
	left, right := 0, len(nums)-1

	// 如果数组没有旋转，直接返回0
	if nums[left] <= nums[right] {
		return 0
	}

	for left <= right {
		mid := left + (right-left)/2

		// 通过比较中间元素与左右相邻元素的关系来判断是否为旋转点
		// 检查mid是否是旋转点
		if mid > 0 && nums[mid-1] < nums[mid] {
			return mid
		}

		// 检查mid+1是否是旋转点
		if mid < len(nums)-1 && nums[mid] > nums[mid+1] {
			return mid + 1
		}

		// 决定搜索方向
		if nums[mid] >= nums[left] {
			// 旋转点在右侧
			left = mid + 1
		} else {
			// 旋转点在左侧
			right = mid - 1
		}
	}

	return 0
}

// 标准的二分查找
func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findPivot1(nums []int) int {
	left, right := 0, len(nums)-1
	// 当 left == right 时，区间里只剩一个元素，必然是最小值（旋转点），所以循环可以停
	// 只要 left < right，说明区间里至少有两个元素还没确定最小值在哪，要继续二分
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
