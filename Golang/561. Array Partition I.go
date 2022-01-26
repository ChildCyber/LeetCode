package leetcode

import "sort"

// 数组拆分 I
// https://leetcode-cn.com/problems/array-partition-i/
// 排序+贪心
func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}
	return
}
