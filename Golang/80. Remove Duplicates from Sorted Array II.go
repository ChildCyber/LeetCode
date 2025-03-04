package leetcode

// 删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/

// 快慢指针
func removeDuplicatesII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	count := 1 // 第一个元素已经出现1次
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[fast-1] {
			count++
		} else {
			count = 1 // 新元素，重置计数
		}

		if count <= 2 {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

func removeDuplicatesIIOptimized(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	slow := 2 // 从第三个位置开始检查

	for fast := 2; fast < len(nums); fast++ {
		// 关键判断：当前元素是否与慢指针前2个位置的元素相同
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}
