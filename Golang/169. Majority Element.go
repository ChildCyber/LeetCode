package leetcode

// 多数元素
// https://leetcode-cn.com/problems/majority-element/
func majorityElement(nums []int) int { // 时间复杂度 O(n) 空间复杂度 O(n)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 {
			return v
		}
	}
	return 0
}

// Boyer-Moore 投票算法 摩尔投票法
// 随机化
