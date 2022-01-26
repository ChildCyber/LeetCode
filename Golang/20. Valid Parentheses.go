package leetcode

// 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/
// 左括号入栈，右括号弹栈
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0)
	for _, v := range s { // string直接迭代是rune类型
		if v == '(' || v == '{' || v == '[' { // 左括号入栈
			stack = append(stack, v)
		} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' ||
			v == ']' && len(stack) > 0 && stack[len(stack)-1] == '[' ||
			v == '}' && len(stack) > 0 && stack[len(stack)-1] == '{' { // 右括号弹栈
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	// 栈是否为空
	return len(stack) == 0
}
