package leetcode

// 乘积最大子数组
// https://leetcode-cn.com/problems/maximum-product-subarray/
// 暴力
func maxProductBrute(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	// 遍历所有可能的子数组起始位置
	for i := 0; i < len(nums); i++ {
		currentProd := 1
		// 遍历所有可能的结束位置
		for j := i; j < len(nums); j++ {
			currentProd *= nums[j]
			if currentProd > maxProd {
				maxProd = currentProd
			}
		}
	}

	return maxProd
}

// 动态规划
func maxProduct(nums []int) int {
	// 维护两个变量，同时跟踪以每个位置结尾的最大乘积和最小乘积
	// 当前的最大乘积可能来自：
	//   当前数字本身
	//   前一个最大乘积 × 当前数字（如果当前数字是正数）
	//   前一个最小乘积 × 当前数字（如果当前数字是负数）
	minimum, maximum, ans := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 { // 乘以一个负数，当前的最大值就变成最小值，而最小值则变成最大值了
			maximum, minimum = minimum, maximum
		}
		// 更新最大值：可能是当前数本身，或者当前数乘以前面的最大值/最小值
		maximum = max(nums[i], maximum*nums[i])
		// 更新最小值：同样考虑三种可能性
		minimum = min(nums[i], minimum*nums[i])
		// 更新全局结果
		ans = max(ans, maximum)
	}
	return ans
}

// 前后两次遍历法
// 最大乘积子数组要么从数组开头开始，要么在数组结尾结束
// 分别从左到右和从右到左遍历一次，计算累积乘积
func maxProductTwo(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProd := nums[0]
	currentProd := 1

	// 从左到右遍历：捕获所有以数组开头为起点的最大乘积子数组
	for i := 0; i < len(nums); i++ {
		currentProd *= nums[i]
		maxProd = max(maxProd, currentProd)
		if currentProd == 0 {
			currentProd = 1 // 重置
		}
	}

	currentProd = 1
	// 从右到左遍历：捕获所有以数组结尾为终点的最大乘积子数组
	for i := len(nums) - 1; i >= 0; i-- {
		currentProd *= nums[i]
		maxProd = max(maxProd, currentProd)
		if currentProd == 0 {
			currentProd = 1 // 重置
		}
	}

	return maxProd
}
