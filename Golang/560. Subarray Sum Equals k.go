package leetcode

// 和为K的子数组
// https://leetcode-cn.com/problems/subarray-sum-equals-k

// 暴力
// 时间复杂度：O(N^3)
// 枚举nums中所有子数组，统计其中和为k的子数组个数
func subarraySumForce(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ { // 固定起点start
		for end := start; end < len(nums); end++ { // 固定终点end
			sum := 0
			for i := start; i <= end; i++ { // nums[start:end+1]范围内和是否等于k
				sum += nums[i]
				if sum == k {
					count++
				}
			}
		}
	}
	return count
}

// 枚举
// 时间复杂度：O(n^2)
// 空间复杂度：O(1)
func subarraySumForce1(nums []int, k int) int {
	// 固定起点，即先固定左边界，然后枚举右边界
	count := 0                                   // 记录和为k的子数组个数
	for start := 0; start < len(nums); start++ { // 遍历nums，考虑每个位置以start作为子数组结尾的情况
		sum := 0                                   // 记录当前子数组的和
		for end := start; end < len(nums); end++ { // 考虑所有0到start位置作为子数组的起点，end作为终点，nums[start:end+1]范围内和是否等于k
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和+哈希表优化
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func subarraySum(nums []int, k int) int {
	/*
		问题转换：子数组和问题 → 两数之差问题
		前缀和数组：第i个元素表示从数组开头到第i个元素（包括第i个元素）的和
		子数组[i, j]的和 = prefix[j] - prefix[i-1]（前缀和[j] 表示从开始到j的所有元素和，前缀和[i] 表示从开始到i-1的所有元素和）
		要找的是 prefix[j] - prefix[i-1] = k，等价于 prefix[i-1] = prefix[j] - k
		这意味着：对于每个位置j，需要找在j之前有多少个位置i，使得prefix[i] = prefix[j] - k
	*/

	count := 0                      // 和为 k 的子数组数量
	currentSum := 0                 // 存储遍历过程中的 当前 前缀和
	prefixSumCount := map[int]int{} // 记录每个前缀和出现的次数，{前缀和: 出现次数}
	prefixSumCount[0] = 1           // 子数组[0,j]的和 = prefix[j] - prefix[-1]，但prefix[-1]不存在，约定prefix[-1] = 0，所以子数组[0,j]的和 = prefix[j] - 0 = prefix[j]；这样可以方便地处理那些从数组第一个元素开始就满足条件的子数组
	// 遍历数组中的每一个数字
	for i := 0; i < len(nums); i++ {
		// 更新当前的 前缀和
		currentSum += nums[i]

		// 要寻找的目标是 oldSum = currentSum - k
		// 如果在之前的位置，存在一个前缀和等于 oldSum，
		// 那么从那个位置之后到当前位置的子数组的和就等于 k。
		need := currentSum - k

		// 在哈希表中查找 need 出现过多少次，并累加到 count 中
		if times, ok := prefixSumCount[need]; ok {
			count += times
		}

		// 将当前的前缀和存入哈希表中，并更新其出现次数
		prefixSumCount[currentSum] += 1
	}
	return count
}
