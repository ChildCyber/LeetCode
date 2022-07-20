package leetcode

// 移除元素
// https://leetcode-cn.com/problems/remove-element/
// 类似283、26

// 双指针
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != val {
			if right != left {
				nums[right], nums[left] = nums[left], nums[right] // 交换
			}
		}
		left++
	}
	return left
}

// 双指针优化
func removeElementP(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		// 左指针与val比较，相同和右指针交换，右指针变为val，同时左移，重复比较左指针和val
		if nums[left] == val { // 相等，右指针左移
			nums[left], nums[right-1] = nums[right-1], nums[left] // 交换
			right--
		} else { // 不等，左指针右移
			left++
		}
	}
	return left
}
