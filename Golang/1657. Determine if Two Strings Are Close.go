package leetcode

// 确定两个字符串是否接近
// https://leetcode.cn/problems/determine-if-two-strings-are-close/

// 哈希表
func closeStrings(word1 string, word2 string) bool {
	// 第一步：检查长度是否相同
	if len(word1) != len(word2) {
		return false
	}

	// 第二步：统计字符频率
	freq1 := make([]int, 26)
	freq2 := make([]int, 26)

	for _, ch := range word1 {
		freq1[ch-'a']++
	}
	for _, ch := range word2 {
		freq2[ch-'a']++
	}

	// 第三步：检查字符集合是否相同
	for i := 0; i < 26; i++ {
		// 如果一个字符在一个字符串中出现而在另一个中不出现
		if (freq1[i] == 0 && freq2[i] != 0) || (freq1[i] != 0 && freq2[i] == 0) {
			return false
		}
	}

	// 第四步：检查频率分布是否相同
	// 统计频率的出现次数
	freqCount1 := make(map[int]int)
	freqCount2 := make(map[int]int)

	for i := 0; i < 26; i++ {
		if freq1[i] > 0 {
			freqCount1[freq1[i]]++
		}
		if freq2[i] > 0 {
			freqCount2[freq2[i]]++
		}
	}

	// 比较频率分布
	for freq, count := range freqCount1 {
		if freqCount2[freq] != count {
			return false
		}
	}

	return true
}
