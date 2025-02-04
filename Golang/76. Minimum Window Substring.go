package leetcode

import "math"

// 最小覆盖子串
// https://leetcode-cn.com/problems/minimum-window-substring

// 暴力
// 思路：枚举子串的起点 i，然后往后扩展 j，判断 [i..j] 是否包含了 t 的所有字符
func minWindowBruteForce(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// 统计 t 中每个字符的需求
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	res := ""
	minLen := math.MaxInt32

	// 枚举所有起点
	for i := 0; i < len(s); i++ {
		// 统计当前子串的频次
		window := make(map[byte]int)

		// 从 i 往后扩展
		for j := i; j < len(s); j++ {
			window[s[j]]++

			// 检查当前子串是否覆盖 t
			if contains(window, need) {
				if j-i+1 < minLen {
					minLen = j - i + 1
					res = s[i : j+1]
				}
				break // 子串已经覆盖了，没必要再扩展
			}
		}
	}

	return res
}

// 判断 window 是否覆盖 need
func contains(window, need map[byte]int) bool {
	for k, v := range need {
		if window[k] < v {
			return false
		}
	}
	return true
}

// 滑动窗口+哈希表
// 时间复杂度：O(n)
// 空间复杂度：O(k)，k为字符集大小
func minWindow(s string, t string) string {
	// 思路：
	// 用两个指针left和right在字符串上定义一个窗口
	// 不断扩展右指针，直到窗口包含了t的所有字符
	// 然后收缩左指针，尽量缩小窗口但仍保持包含所有字符
	// 记录最小的满足条件的窗口
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// 1. 统计 t 中每个字符的需求
	need := make(map[byte]int) // 记录目标字符及其所需数量
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 2. 定义窗口
	window := make(map[byte]int) // 记录窗口内的字符情况
	left, right := 0, 0          // left:收缩窗口，优化解；right:扩展窗口，寻找可行解
	valid := 0                   // 记录窗口中满足要求的字符个数，用于快速判断窗口是否满足条件

	// 3. 记录最优解
	start, length := 0, math.MaxInt32

	// 4. 移动右指针
	for right < len(s) {
		c := s[right]
		right++ // 窗口的右边界始终指向下一个待处理的元素

		// 更新窗口数据
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// 5. 当窗口满足条件时，尝试收缩左边
		for valid == len(need) { // 说明窗口已经覆盖 t，开始收缩左指针，更新最优解
			if right-left < length { // 候选答案
				start = left
				length = right - left
			}

			d := s[left]
			left++
			// 只有当被移出的字符 d 是 t 中需要的字符时，才需要去维护 window 和 valid
			if _, ok := need[d]; ok {
				// 如果在移除之前 window[d] 刚好等于 need[d]，说明窗口对字符 d 的需求是“恰好满足”的
				if window[d] == need[d] {
					valid-- // 一旦移除一个 d，其频次就会变小，导致不再满足需求，于是 valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
