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

// 快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0 // 慢指针，指向当前唯一元素的末尾位置
	// 快指针从第2个元素开始遍历
	for fast := 1; fast < len(nums); fast++ { // 快指针表示遍历数组到达的下标位置
		// 当遇到不同元素时，移动慢指针并更新值
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
