package leetcode

// 子数组最大平均数I
// https://leetcode.cn/problems/maximum-average-subarray-i/

// 滑动窗口
func findMaxAverage(nums []int, k int) float64 {
	// 计算第一个窗口的和
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}

	maxSum := currentSum

	// 滑动窗口，从第k个元素开始
	for i := k; i < len(nums); i++ {
		// 减去离开窗口的元素，加上新进入窗口的元素
		currentSum = currentSum - nums[i-k] + nums[i]

		// 更新最大和
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	// 返回最大平均值
	return float64(maxSum) / float64(k)
}
