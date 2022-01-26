package leetcode

// 和为K的子数组
// https://leetcode-cn.com/problems/subarray-sum-equals-k
// 暴力
func subarraySum(nums []int, k int) int {
	count := 0                                   // 记录和为k的子数组个数
	for start := 0; start < len(nums); start++ { // 遍历nums，考虑每个位置以start作为子数组结尾的情况；从前向后
		sum := 0                            // 记录当前子数组的和
		for end := start; end >= 0; end-- { // 考虑所有0到start位置作为子数组的开头；从后向前
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
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}
