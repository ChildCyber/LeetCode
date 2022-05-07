package leetcode

// 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/
// 回溯
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	findGenerateParenthesis(n, n, "", &ans) // 从n开始，不是从0开始
	return ans
}

func findGenerateParenthesis(lindex, rindex int, str string, ans *[]string) {
	// 不需要判断括号是否匹配。在 DFS 回溯的过程中，会让 ( 和 ) 成对的匹配上
	// 只在序列仍然保持有效时才添加，通过跟踪到目前为止放置的左括号和右括号的数目来做到这一点
	if lindex == 0 && rindex == 0 {
		*ans = append(*ans, str)
		return
	}
	// 如果左括号数量不大于 n，可以放一个左括号。如果右括号数量小于左括号的数量，可以放一个右括号。
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", ans)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", ans)
	}
}

// 暴力 + 递归
// 生成所有序列，然后检查每一个序列是否有效
// 时间复杂度O(2^2n)，空间复杂度O(n)
func generateParenthesisForce(n int) []string {
	if n == 0 {
		return []string{}
	}
	ans := []string{}
	generateAll22("", 0, &ans)
	return ans
}

// 生成所有序列
// 递归：生成长度为2n的字符串，等价于生成第一个字符 加上 剩下为2n-1个字符
func generateAll22(current string, pos int, ans *[]string) {
	if len(current) == pos { // 字符构造完成，判断是否符合条件
		if valid22(current) {
			*ans = append(*ans, current)
		}
	} else { // 递归生成序列
		// 假设当前字符填入左括号或右括号
		current += "("
		generateAll22(current, pos+1, ans)
		current = current[:len(current)-1]
		current += ")"
		generateAll22(current, pos+1, ans)
	}
}

// 验证是否为有效的括号
func valid22(current string) bool {
	balance := 0
	// 自左向右读取字符串时，左括号的个数始终大于右括号的个数
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
