package leetcode

// 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// 回溯
// 时间复杂度：O(4^n)，其中n是输入数字字符串的长度
// 空间复杂度：O(n)，主要用于递归调用栈
func letterCombinationsBT(digits string) []string {
	ans := []string{}
	if digits == "" {
		return ans
	}

	// 存储每个数字对应的所有可能的字母
	var digitMap = map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var path []byte

	var dfs func(pos int)
	dfs = func(pos int) {
		// 如果拼完了所有数字，就加到结果里
		if pos == len(digits) {
			ans = append(ans, string(path))
			return
		}

		letters := digitMap[string(digits[pos])]
		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			dfs(pos + 1)
			path = path[:len(path)-1] // 撤销
		}
	}

	dfs(0)
	return ans
}
