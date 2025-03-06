package leetcode

// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

// 暴力
func lengthOfLongestSubstringBrute(s string) int {
	n := len(s)
	ans := 0

	// 枚举所有子串，判断每个是否含重复字符
	for i := 0; i < n; i++ {
		seen := make(map[byte]bool)
		for j := i; j < n; j++ {
			if seen[s[j]] {
				break
			}
			seen[s[j]] = true
			if j-i+1 > ans {
				ans = j - i + 1
			}
		}
	}

	return ans
}

// 滑动窗口+哈希表
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	/*
		滑动窗口的右边界不断的右移，只要没有重复的字符，就持续向右扩大窗口边界。
		一旦出现了重复字符，就需要缩小左边界，直到重复的字符移出了左边界，然后继续移动滑动窗口的右边界。
		以此类推，每次移动需要计算当前⻓度，并判断是否需要更新最大⻓度，最终最大的值就是题目中的所求。
	*/
	var bitSet [256]uint8       // 位图，记录每个字符是否出现过
	ans, left, right := 0, 0, 0 // 结果，左指针，右指针

	// 窗口定义：左闭右开区间
	for left < len(s) {
		// 没有重复的字符，向右扩大窗口边界
		if right < len(s) && bitSet[s[right]] == 0 { // 判断右窗口是否越界，字符是否首次出现
			bitSet[s[right]] = 1 // 该字符标记为已出现
			right++              // 扩大右窗口
		} else { // 出现了重复字符，缩小左边界
			bitSet[s[left]] = 0 // 该字符标记为未出现
			left++              // 扩大左窗口
		}
		ans = max(ans, right-left) // 存储最大结果值
	}

	return ans
}

// 跳过重复字符
func lengthOfLongestSubstringOptimized(s string) int {
	last := make(map[byte]int) // 记录字符上次出现的位置
	ans := 0
	left := 0 // 窗口当前第一个字符的索引

	// right：下一个准备加入窗口的字符索引（所以它还不在窗口里）
	for right := 0; right < len(s); right++ { // [left, right) = [0, 0) 表示一个空窗口。窗口内的字符是 s[left:right]，即 s[0:0]，长度为 0
		ch := s[right]
		if idx, ok := last[ch]; ok && idx >= left {
			// 如果ch在窗口中出现过，移动左边界
			left = idx + 1
		}
		last[ch] = right
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
