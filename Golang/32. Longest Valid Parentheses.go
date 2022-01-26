package leetcode

// hard
// 最长有效括号
// https://leetcode-cn.com/problems/longest-valid-parentheses/
// 栈
func longestValidParentheses(s string) int {
	maxAns := 0
	stack := []int{}
	// 如果一开始栈为空，第一个字符为左括号的时候会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，在一开始的时候往栈中放入一个值为 −1 的元素
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' { // 下标入栈
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 先弹出栈顶元素表示匹配了当前右括号
			if len(stack) == 0 {         // 如果栈为空，说明当前的右括号为没有被匹配的右括号，将其下标放入栈中来更新「最后一个没有被匹配的右括号的下标」
				stack = append(stack, i)
			} else { // 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}
