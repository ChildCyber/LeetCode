package leetcode

import "sort"

// 咒语和药水的成功对数
// https://leetcode.cn/problems/successful-pairs-of-spells-and-potions/

// 排序+二分
func successfulPairs(spells []int, potions []int, success int64) []int {
	// 对药水数组进行排序
	sort.Ints(potions)
	n := len(spells)
	m := len(potions)
	ans := make([]int, n)

	for i := 0; i < n; i++ {
		spell := spells[i]

		// 计算需要的最小药水强度，向上取整
		minPotion := (int(success) + spell - 1) / spell

		// 二分查找第一个 >= minPotion 的位置
		left, right := 0, m
		for left < right {
			mid := left + (right-left)/2
			if potions[mid] < minPotion {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 计算成功配对数
		ans[i] = m - left
	}

	return ans
}
