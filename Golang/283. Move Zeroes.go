package leetcode

// 移动零
// https://leetcode-cn.com/problems/move-zeroes/
// 双指针
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++ { // 一次遍历，类似快排
		if nums[i] != 0 { // 用 i,j 标记 0 和非 0 的元素，然后相互交换
			if i != j { // i为已遍历位置，j为以处理区间末尾
				nums[i], nums[j] = nums[j], nums[i]
				j++
			} else {
				j++
			}
		}
	}
}

func moveZeroes1(nums []int) {
	// 左指针指向当前已经处理好的序列的尾部，右指针指向待处理序列的头部
	// 右指针不断向右移动，每次右指针指向非零数，则将左右指针对应的数交换，同时左指针右移
	// 左指针左边均为非零数； 右指针左边直到左指针处均为零
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
