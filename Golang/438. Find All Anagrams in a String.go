package leetcode

// 找到字符串中所有字母异位词
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/

// 滑动窗口
// 时间复杂度：O(n)，其中 n 是字符串 s 的长度
// 空间复杂度：O(1)
func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	ans := []int{}
	pCount := [26]int{}
	sCount := [26]int{}

	// 统计 p 中每个字符的出现频率
	for i := 0; i < len(p); i++ {
		pCount[p[i]-'a']++
	}

	// 初始化滑动窗口
	left := 0
	for right := 0; right < len(s); right++ {
		// 将右指针指向的字符加入窗口
		sCount[s[right]-'a']++

		// 当窗口大小等于 p 的长度时
		if right-left+1 == len(p) {
			// 检查当前窗口是否与 p 匹配
			if pCount == sCount {
				ans = append(ans, left)
			}

			// 移动左指针，移除左指针指向的字符
			sCount[s[left]-'a']--
			left++
		}
	}

	return ans
}

// 滑动窗口-窗口大小固定
func findAnagrams1(s, p string) (ans []int) {
	// 字符串 p 的异位词的长度一定与字符串 p 的长度相同，可以在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量
	// 当窗口中每种字母的数量与字符串 p 中每种字母的数量相同时，则说明当前窗口为字符串 p 的异位词
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	// 存放字符串中字母出现的词频
	var sCount, pCount [26]int
	// 统计 p 的字符频率和 s 的前pLen个字符频率
	for i := 0; i < pLen; i++ {
		pCount[p[i]-'a']++
		sCount[s[i]-'a']++
	}

	// 判断放置处是否有异位词(在放置时只需判断一次)
	if sCount == pCount {
		ans = append(ans, 0)
	}

	// 滑动窗口遍历剩余部分
	for i := 0; i < sLen-pLen; i++ {
		// 移除左边界字符
		sCount[s[i]-'a']--
		// 添加右边界字符
		sCount[s[i+pLen]-'a']++

		// 判断滑动后处，是否有异位词
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}

	return
}
