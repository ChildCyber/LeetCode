package leetcode

// 寻找数组的中心下标
// https://leetcode.cn/problems/find-pivot-index/

// 暴力
// 对每个i分别求左边和右边再比较
func pivotIndexBrute(nums []int) int {
	n := len(nums)

	// 遍历每个可能的中点
	for i := 0; i < n; i++ {
		leftSum := 0
		rightSum := 0

		// 计算左侧和
		for j := 0; j < i; j++ {
			leftSum += nums[j]
		}

		// 计算右侧和
		for j := i + 1; j < n; j++ {
			rightSum += nums[j]
		}

		// 检查是否满足条件
		if leftSum == rightSum {
			return i
		}
	}

	return -1
}

// 前缀和-一次遍历
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	left := 0 // 当前位置左侧的和
	for i, v := range nums {
		if left == total-left-v { // 当前位置右侧的和：total - left - nums[i]
			return i
		}
		left += v
	}
	return -1
}
