package leetcode

// 移动零
// https://leetcode-cn.com/problems/move-zeroes/
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
