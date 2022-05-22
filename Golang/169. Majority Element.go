package leetcode

// 多数元素
// https://leetcode-cn.com/problems/majority-element/

// 哈希表
// 时间复杂度 O(n) 空间复杂度 O(n)
func majorityElement(nums []int) int {
	// 存储每个元素以及出现的次数。键表示一个元素，值表示该元素出现的次数
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
		if m[v] > len(nums)/2 { // 遍历数组时直接判断，省去了对哈希表的遍历
			return v
		}
	}
	return 0
}

// Boyer-Moore 投票算法 摩尔投票法
// 随机化
