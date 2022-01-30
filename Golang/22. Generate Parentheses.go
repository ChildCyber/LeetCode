package leetcode

// 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/
// 回溯
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

// 暴力 + 递归
func generateParenthesisForce(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	generateAll22("", 0, &res)
	return res
}

func generateAll22(current string, pos int, res *[]string) {
	if len(current) == pos {
		if valid22(current) {
			*res = append(*res, current)
		}
	} else {
		current += "("
		generateAll22(current, pos+1, res)
		current = current[:len(current)-1]
		current += ")"
		generateAll22(current, pos+1, res)
	}
}

func valid22(current string) bool {
	balance := 0
	for i := 0; i < len(current); i++ {
		if current[i] == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
