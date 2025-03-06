package leetcode

// 两数之和 II - 输入有序数组
// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/

// 双指针
func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil // 题目保证有解，这里不会执行到
}
