package leetcode

// 和为K的子数组
// https://leetcode-cn.com/problems/subarray-sum-equals-k
// 暴力
// 枚举nums中所有子数组，统计其中和为k的子数组个数
func subarraySumForce(nums []int, k int) int {
	count := 0
	for start := 0; start < len(nums); start++ {
		for end := start; end < len(nums); end++ {
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
func subarraySum(nums []int, k int) int {
	// 固定了起点，即先固定左边界，然后枚举右边界
	count := 0                                   // 记录和为k的子数组个数
	for start := 0; start < len(nums); start++ { // 遍历nums，考虑每个位置以start作为子数组结尾的情况
		sum := 0                                   // 记录当前子数组的和
		for end := start; end < len(nums); end++ { // 考虑所有0到start位置作为子数组的开头
			sum += nums[end]
			if sum == k {
				count++
			}
		}
	}
	return count
}

// 前缀和+哈希表优化
func subarraySum1(nums []int, k int) int {
	count, pre := 0, 0
	// key：前缀和
	// value：key 对应的前缀和的个数
	preSumFreq := map[int]int{}
	// 对于下标为 0 的元素，前缀和为 0，个数为 1
	preSumFreq[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		// 先获得前缀和为 pre - k 的个数，加到计数变量里
		if _, ok := preSumFreq[pre-k]; ok {
			count += preSumFreq[pre-k]
		}
		// 然后维护 preSumFreq 的定义
		preSumFreq[pre] += 1
	}
	return count
}
