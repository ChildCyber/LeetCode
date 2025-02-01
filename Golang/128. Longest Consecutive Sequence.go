package leetcode

// 最长连续序列
// https://leetcode-cn.com/problems/longest-consecutive-sequence/
// 要求时间复杂度为 O(n)
// 哈希表
func longestConsecutive(nums []int) int {
	longestStreak := 0 // 记录连续的最长序列的长度
	// 遍历初始 nums, 使用哈希表存储出现过的元素
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	// 枚举数组中的每个数 x，以其为起点，不断尝试匹配 x+1,x+2,⋯x+1, x+2, ⋯ 是否存在
	// 假设最长匹配到了 x+y，那么以 x 为起点的最长连续序列即为 x,x+1,x+2,⋯,x+y，其长度为 y+1，不断枚举并更新答案即可
	for num := range numSet { // 遍历numSet，num为map中的key
		// 跳过逻辑：只有当num是序列起点时才开始计数，num-1不存在则证明num为连续序列的起点
		if !numSet[num-1] { // 枚举的数 num 在数组中一定不存在前驱数 num−1，每次在哈希表中检查是否存在 num−1 即能判断是否需要跳过了
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] { // 获取连续序列的最长长度
				currentNum++
				currentStreak++
			}

			if longestStreak < currentStreak { // 更新最长序列长度
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}
