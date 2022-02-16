package leetcode

// 搜索旋转排序数组
// https://leetcode-cn.com/problems/search-in-rotated-sorted-array
// 算法时间复杂度必须是 O(log n)
// 二分查找
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环。
func search33(nums []int, target int) int {
	//数组前面一段是数值比较大的数，后面一段是数值偏小的数。
	//如果 mid 落在了前一段数值比较大的区间内了，那么一定有nums[mid] > nums[low] ,
	//如果是落在后面一段数值比较小的区间内, nums[mid] ≤ nums[low] 。
	//如果 mid 落在了后一段数值比较小的区间内了,那么一定有 nums[mid] < nums[high] ,如果是落在前面一段
	//数值比较大的区间内, nums[mid] ≤ nums[high] 。
	//还有 nums[low] == nums[mid] 和 nums[high] == nums[mid] 的情况，单独处理即可。
	//最后找到则输出 mid,没有找到则输出 -1 。
	if len(nums) == 0 {
		return -1
	}
	// 二分
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1 // 二分位置的下标
		// 查看当前 mid 为分割位置分割出来的两个部分 [l, mid] 和 [mid + 1, r] 哪个部分是有序的，并根据有序的那个部分确定该如何改变二分查找的上下界，因为能够根据有序的那部分判断出 target 在不在这个部分：
		//    如果 [l, mid - 1] 是有序数组，且 target 的大小满足 [nums[l],nums[mid]]，则应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找。
		//    如果 [mid, r] 是有序数组，且 target 的大小满足 [nums[mid+1],nums[r]]，则应该将搜索范围缩小至 [mid + 1, r]，否则在 [l, mid - 1] 中寻找。
		if nums[mid] == target { // mid即为target
			return mid
		} else if nums[mid] > nums[low] { // 在数值大的一部分区间里
			// [l, mid - 1] 是有序数组
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[high] { // 在数值小的一部分区间里
			if nums[mid] < target && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		} else {
			if nums[low] == nums[mid] {
				low++
			} else {
				high--
			}
		}
	}
	return -1
}

func search33_1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	// 重要变量：low, high, mid
	low, high := 0, len(nums)-1 // 首，尾
	for low <= high {
		mid := (low + high) / 2
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
