package leetcode

// 删除有序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 将删除元素移动到数组后面的空间，返回数组实际剩余的元素个数
// 不要使用额外的空间，必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成

// 双指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	last, finder := 0, 0     // last不重复数字数组最后一个元素的下标，finder遍历中的有序数组的下标
	for last < len(nums)-1 { // 保证last不越界
		for nums[finder] == nums[last] { // 找到当前数字，最后一个不重复元素的下标；当相邻元素不同时，last==finder
			finder++
			if finder == len(nums) { // 找到有序数组的末尾，保证finder不越界
				return last + 1
			}
		}
		nums[last+1] = nums[finder] // 如果不同，两个元素间隔为1，last+1==finder；注意这里是赋值，不是交换
		last++
	}
	return last + 1
}
