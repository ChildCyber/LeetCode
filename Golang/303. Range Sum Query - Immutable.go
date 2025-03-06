package leetcode

// 区域和检索 - 数组不可变
// https://leetcode.cn/problems/range-sum-query-immutable/

// 前缀和
type NumArray struct {
	prefixSum []int
}

func Constructor303(nums []int) NumArray {
	n := len(nums)
	// 创建前缀和数组，长度比原数组多1，prefixSum[0] = 0
	prefixSum := make([]int, n+1)

	// 计算前缀和
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	return NumArray{prefixSum: prefixSum}
}

func (this *NumArray) SumRange(left int, right int) int {
	// 区间和 = 前缀和[right+1] - 前缀和[left]
	return this.prefixSum[right+1] - this.prefixSum[left]
}
