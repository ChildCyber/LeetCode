package leetcode

// 移动零
// https://leetcode-cn.com/problems/move-zeroes/

// 双指针
// 左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部
// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移
// 左指针左边均为非零数；右指针左边直到左指针处均为零
func moveZeroes1(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
