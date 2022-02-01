package leetcode

// 寻找峰值
// https://leetcode-cn.com/problems/find-peak-element/
// 间复杂度为 O(log n)
// 寻找最大值，数组 nums 中最大值两侧的元素一定严格小于最大值本身。因此，最大值所在的位置就是一个可行的峰值位置。
func findPeakElement(nums []int) (idx int) {
	for i, v := range nums {
		if v > nums[idx] {
			idx = i
		}
	}
	return
}

// 二分，迭代爬坡
// 从一个位置开始，不断地向高处走，那么最终一定可以到达一个峰值位置。
