package leetcode

// 电话号码的字母组合
// https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
// 回溯
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

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	// index：取到digits中的第几个字母
	backtrack17(digits, 0, "")
	return combinations
}

func backtrack17(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		for i := 0; i < len(letters); i++ {
			// 由于每个数字对应的每个字母都可能进入字母组合，因此不存在不可行的解，直接穷举所有的解即可
			backtrack17(digits, index+1, combination+string(letters[i])) // 每次取电话号码的一位数字
		}
	}
}
