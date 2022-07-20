package leetcode

// 删除有序数组中的重复项
// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
// 将删除元素移动到数组后面的空间，返回数组实际剩余的元素个数
// 不要使用额外的空间，必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成

// 暴力
func removeDuplicatesForce(nums []int) int {
	ans := make([]int, len(nums))
	var index, i int

	// 常规情况
	for i = 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			ans[index] = nums[i]
			index++
			ans[index] = nums[i+1]
		}
	}
	// 输入全部的数组都相同
	if index == 0 {
		ans[index] = nums[i]
		return 1
	}
	// 复制回老数组
	for j := 0; j < index+1; j++ {
		nums[j] = ans[j]
	}
	return index + 1
}

// 双指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	last, finder := 0, 0     // last：不重复数字数组最后一个元素的下标，finder：遍历中的有序数组的下标
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

func removeDuplicates1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	slow := 1 // 慢指针表示下一个不同元素要填入的下标位置
	// 当nums的长度大于0，数组中至少包含一个元素，在删除重复元素之后也至少剩下一个元素，因此nums[0]保持原状即可，从下标1开始删除重复元素
	for fast := 1; fast < n; fast++ { // 快指针表示遍历数组到达的下标位置
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
