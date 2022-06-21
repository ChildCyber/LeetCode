package leetcode

// 两个数组的交集 II
// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
// 返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）

// 哈希表
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	// 用哈希表存储每个数字出现的次数。对于一个数字，其在交集中出现的次数等于该数字在两个数组中出现次数的最小值
	m := map[int]int{}
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 { // 当前元素在nums1中出现过
			ans = append(ans, n)
			m[n]-- // 减少该数字出现的次数
		}
	}
	return ans
}
