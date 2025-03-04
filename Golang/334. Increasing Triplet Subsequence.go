package leetcode

import "math"

// 递增的三元子序列
// https://leetcode.cn/problems/increasing-triplet-subsequence/

// 双向遍历
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func increasingTriplet(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	leftMin := make([]int, n) // 记录每个元素左边的最小值
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ { // 从左到右遍历
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	rightMax := make([]int, n) // 记录每个元素右边的最大值
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- { // 从右到左遍历
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	// 检查三元组，是否存在leftMin[i-1] < nums[i] < rightMax[i+1]
	for i := 1; i < n-1; i++ {
		if nums[i] > leftMin[i-1] && nums[i] < rightMax[i+1] { // 当前元素左边存在一个元素小于nums[i]，并且右边存在一个元素大于nums[i]
			return true
		}
	}

	return false
}

// 贪心
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func increasingTriplet1(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	first, second := nums[0], math.MaxInt32 // 分别表示递增的三元子序列中的第一个数和第二个数，任何时候都有first<second
	for i := 1; i < n; i++ {                // 从左到右遍历数组
		num := nums[i]
		if num <= first {
			// 更新最小值 => nums[i}<first<second
			first = num
		} else if num <= second {
			// 更新第二小值 => first<nums[i]<second
			second = num
		} else {
			// 找到比second更大的数 => first<second<nums[i]
			return true
		}
	}

	return false
}
