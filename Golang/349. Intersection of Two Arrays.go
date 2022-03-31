package leetcode

// 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	var res []int
	for _, n := range nums1 {
		m[n] = true
	}
	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}
