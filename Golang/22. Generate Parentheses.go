package leetcode

// 括号生成
// https://leetcode-cn.com/problems/generate-parentheses/

// 回溯-使用字节切片（显式回溯）
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	path := make([]byte, 0, 2*n) // 预分配容量，避免频繁扩容

	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		// 左右括号都用完，保存结果
		if left == 0 && right == 0 {
			ans = append(ans, string(path))
			return
		}

		// 还能放左括号
		if left > 0 {
			path = append(path, '(')
			backtrack(left-1, right)
			path = path[:len(path)-1] // 回溯
		}

		// 还能放右括号（右比左多时）
		if right > left {
			path = append(path, ')')
			backtrack(left, right-1)
			path = path[:len(path)-1] // 回溯
		}
	}

	backtrack(n, n)
	return ans
}

// 回溯-使用字符串（隐式回溯，不需要恢复状态）
func generateParenthesis1(n int) []string {
	ans := make([]string, 0)

	var backtrack func(int, int, string)
	backtrack = func(left, right int, path string) {
		if left == 0 && right == 0 {
			ans = append(ans, path)
			return
		}

		if left > 0 {
			backtrack(left-1, right, path+"(")
		}
		if right > left {
			backtrack(left, right-1, path+")")
		}
	}

	backtrack(n, n, "")
	return ans
}

// 暴力+递归
// 时间复杂度：O(2^2n)
// 空间复杂度：O(n)
func generateParenthesisBrute(n int) []string {
	// 思路：生成所有序列，然后检查每一个序列是否有效
	if n == 0 {
		return []string{}
	}

	ans := []string{}
	generateAll22("", 2*n, &ans)
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
		generateAll22(current, pos, ans)
		current = current[:len(current)-1] // 回溯，选择右括号
		current += ")"
		generateAll22(current, pos, ans)
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
