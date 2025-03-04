package leetcode

// 移除元素
// https://leetcode-cn.com/problems/remove-element/
// 类似283、26

// 快慢指针
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return slow
}

// 对撞指针
func removeElementP(nums []int, val int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// 左指针与val比较，相同和右指针交换，右指针变为val，同时左移，重复比较左指针和val
		if nums[left] == val { // 相等，右指针左移
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else { // 不等，左指针右移
			left++
		}
	}
	return left
}
