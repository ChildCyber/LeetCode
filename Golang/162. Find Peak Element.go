package leetcode

// 寻找峰值
// https://leetcode-cn.com/problems/find-peak-element/
// 寻找最大值
func findPeakElement(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}
