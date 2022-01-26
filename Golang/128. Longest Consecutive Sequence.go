package leetcode

// 最长连续序列
// https://leetcode-cn.com/problems/longest-consecutive-sequence/
// 时间复杂度为 O(n)
// 哈希表
func longestConsecutive(nums []int) int {
	// 遍历初始 nums, 使用哈希表存储出现过的元素, 并且去重
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0        // 连续的最长序列的长度
	for num := range numSet { // 遍历numSet，num为map中的key
		// 不在集合中，更新序列长度；在集合中，跳过
		if !numSet[num-1] { // 在数组中不存在前驱数 num−1
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] { // 集合中存在
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
