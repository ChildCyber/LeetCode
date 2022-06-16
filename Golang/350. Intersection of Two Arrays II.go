package leetcode

// 两个数组的交集 II
// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
// 哈希表
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	m := map[int]int{}
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if m[n] > 0 {
			ans = append(ans, n)
			m[n]--
		}
	}
	return ans
}
