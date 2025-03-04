package leetcode

// 定长子串中元音的最大数目
// https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/

// 滑动窗口
func maxVowels(s string, k int) int {
	// 创建元音集合
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

	// 计算第一个窗口的元音数量
	currentCount := 0
	for i := 0; i < k; i++ {
		if vowels[s[i]] {
			currentCount++
		}
	}

	maxCount := currentCount

	// 滑动窗口
	for i := k; i < len(s); i++ {
		// 减去离开窗口的字符的元音值
		if vowels[s[i-k]] {
			currentCount--
		}

		// 加上新进入窗口的字符的元音值
		if vowels[s[i]] {
			currentCount++
		}

		// 更新最大值
		if currentCount > maxCount {
			maxCount = currentCount
		}

		// 如果已经达到最大值，提前返回
		if maxCount == k {
			return maxCount
		}
	}

	return maxCount
}
