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

	// 枚举数组中的每个数 x，考虑以其为起点，不断尝试匹配 x+1,x+2,⋯x+1, x+2, ⋯ 是否存在，
	// 假设最长匹配到了 x+y，那么以 x 为起点的最长连续序列即为 x,x+1,x+2,⋯,x+y，其长度为 y+1，不断枚举并更新答案即可。
	longestStreak := 0        // 连续的最长序列的长度
	for num := range numSet { // 遍历numSet，num为map中的key
		// 不在集合中，更新序列长度；在集合中，跳过
		// 枚举的数 x 一定是在数组中不存在前驱数 x−1，每次在哈希表中检查是否存在 x−1 即能判断是否需要跳过了
		if !numSet[num-1] { // 在数组中不存在前驱数 num−1
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] { // 查找 num+n
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
