package leetcode

// 独一无二的出现次数
// https://leetcode.cn/problems/unique-number-of-occurrences/

// 哈希表
func uniqueOccurrences(arr []int) bool {
	// 第一步：统计每个数字的出现次数
	freqMap := make(map[int]int)
	for _, num := range arr {
		freqMap[num]++
	}

	// 第二步：检查出现次数是否唯一
	freqSet := make(map[int]bool)
	for _, count := range freqMap {
		if freqSet[count] {
			return false // 发现重复的出现次数
		}
		freqSet[count] = true
	}

	return true
}
