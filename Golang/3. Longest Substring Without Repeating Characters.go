package leetcode

// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	// 滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。一旦出现了重复字符，就需要缩
	// 小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。以此类推，每次移动需要计算当前⻓
	// 度，并判断是否需要更新最大⻓度，最终最大的值就是题目中的所求。
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]bool // 位图，记录每个字符是否出现过
	result, left, right := 0, 0, 0
	for left < len(s) {
		// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将 X 标记为 false
		if bitSet[s[right]] { // 滑动窗口
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}

		if result < right-left { // 存储最大结果值
			result = right - left
		}

		if left+result >= len(s) || right >= len(s) { // 判断是否越界
			break
		}
	}
	return result
}

// 滑动窗口
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]uint8       // 位图，记录每个字符是否出现过
	res, left, right := 0, 0, 0 // 结果，左指针，右指针(都从头开始)
	for left < len(s) {
		// 没有重复的字符，向右扩大窗口边界
		if right < len(s) && bitSet[s[right]] == 0 { // 判断右窗口是否越界，该字符首次出现（重点）
			bitSet[s[right]] = 1 // 该字符标记为已出现
			right++              // 扩大右窗口
		} else { // 出现了重复字符，缩小左边界
			bitSet[s[left]] = 0 // 该字符标记为未出现
			left++              // 扩大左窗口
		}
		res = max(res, right-left) // 存储最大结果值（重点）
	}
	return res
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}

	var bitSet [256]int
	index, ans := 0, 0
	for i := 0; i < len(s); i++ { // 遍历字符串
		if bitSet[s[i]] != 0 {
			index = max(i, bitSet[s[i]])
		}
		// 将左指针i的位置变为 第一次重复 的字符的下一个
		bitSet[s[i]] = i + 1
		ans = max(ans, index-i+1)
	}
	return ans
}
