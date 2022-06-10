package leetcode

// 找到字符串中所有字母异位词
// https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/
// 滑动窗口
func findAnagrams(s, p string) (ans []int) {
	// 字符串 p 的异位词的长度一定与字符串 p 的长度相同，可以在字符串 s 中构造一个长度为与字符串 p 的长度相同的滑动窗口，并在滑动中维护窗口中每种字母的数量
	// 当窗口中每种字母的数量与字符串 p 中每种字母的数量相同时，则说明当前窗口为字符串 p 的异位词
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return
	}

	// 建立两个数组存放字符串中字母出现的词频，并以此作为标准比较
	var sCount, pCount [26]int
	// 当滑动窗口的首位在s[0]处时 （相当于放置滑动窗口进入数组）
	for i, ch := range p {
		sCount[s[i]-'a']++ // s[0:p]判断
		pCount[ch-'a']++
	}
	// 判断放置处是否有异位词     (在放置时只需判断一次)
	if sCount == pCount {
		ans = append(ans, 0)
	}

	// 开始让窗口进行滑动
	for i, ch := range s[:sLen-pLen] { // i是滑动前的首位
		sCount[ch-'a']--        // 将滑动前首位的词频删去
		sCount[s[i+pLen]-'a']++ // 增加滑动后最后一位的词频（以此达到滑动的效果）
		// 判断滑动后处，是否有异位词
		if sCount == pCount {
			ans = append(ans, i+1)
		}
	}
	return
}
