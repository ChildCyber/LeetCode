package leetcode

// 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// 问题本质：组合生成问题

// 回溯
// 时间复杂度：O(4^n)，其中n是输入数字字符串的长度
// 空间复杂度：O(n)，主要用于递归调用栈
func letterCombinations(digits string) []string {
	ans := []string{}
	if digits == "" {
		return ans
	}

	// 存储每个数字对应的所有可能的字母
	var digitMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var path []byte

	var backtrack func(pos int)
	backtrack = func(pos int) {
		// 如果拼完了所有数字，就加到结果里
		if pos == len(digits) {
			ans = append(ans, string(path))
			return
		}

		letters := digitMap[digits[pos]]
		// 从letters里开始选择
		for i := 0; i < len(letters); i++ {
			// 做选择
			path = append(path, letters[i])
			// 递归
			backtrack(pos + 1)
			// 撤销选择（回溯）
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return ans
}
