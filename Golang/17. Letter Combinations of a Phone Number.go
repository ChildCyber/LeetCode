package leetcode

// 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/

// 存储每个数字对应的所有可能的字母
var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 已有的字母排列
var combinations []string

// DFS
// 由于每个数字对应的每个字母都可能进入字母组合，因此不存在不可行的解，直接穷举所有的解即可
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{} // 表示已有的字母排列
	// index：取到digits中的第几个字母
	backtrack17(digits, 0, "") // combinations初始为空
	return combinations
}

func backtrack17(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}

	// 每次取电话号码的一位数字
	digit := string(digits[index])
	// 从哈希表中获得该数字对应的所有可能的字母
	letters := phoneMap[digit]
	// 并将其中的一个字母插入到已有的字母排列后面
	for i := 0; i < len(letters); i++ {
		// 由于每个数字对应的每个字母都可能进入字母组合，因此不存在不可行的解，直接穷举所有的解即可
		backtrack17(digits, index+1, combination+string(letters[i])) // 然后继续处理电话号码的后一位数字，直到处理完电话号码中的所有数字，即得到一个完整的字母排列
	}
	return
}

// 回溯
func letterCombinationsBT(digits string) []string {
	ans := []string{}
	if digits == "" {
		return ans
	}

	var letterFunc func(string, string)
	letterFunc = func(comb string, s string) {
		// comb表示已有的字母排列
		// s避免对原字符串的修改
		if s == "" {
			ans = append(ans, comb)
			return
		}

		k := s[0:1] // 避免类型转换
		letters := phoneMap[k]
		for i := 0; i < len(letters); i++ { // 遍历map中对应key的字符
			// 选择该字符
			comb += string(letters[i]) // 强制类型转化，phoneMap[k][i]为uint8
			letterFunc(comb, s[1:])
			// 回溯，不选择该字符
			comb = comb[0 : len(comb)-1]
		}
	}

	letterFunc("", digits) // comb字符串初始为空
	return ans
}
