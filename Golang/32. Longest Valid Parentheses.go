package leetcode

// hard
// 最长有效括号
// https://leetcode-cn.com/problems/longest-valid-parentheses/
// 栈
func longestValidParentheses(s string) int {
	// 通过栈，在遍历给定字符串的过程中去判断到目前为止扫描的子串的有效性，同时能得到最长有效括号的长度
	maxAns := 0
	stack := []int{}
	// 如果一开始栈为空，第一个字符为左括号的时候会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，在一开始的时候往栈中放入一个值为 −1 的元素
	stack = append(stack, -1)     // 重点，始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」
	for i := 0; i < len(s); i++ { // 遍历字符串
		if s[i] == '(' { // '('做入栈操作
			stack = append(stack, i) // 下标入栈
		} else { // ')'做出栈操作
			stack = stack[:len(stack)-1] // 先弹出栈顶元素，表示匹配了当前右括号
			if len(stack) == 0 {         // 如果栈为空，说明当前的右括号为没有被匹配的右括号，将其下标放入栈中来更新「最后一个没有被匹配的右括号的下标」
				stack = append(stack, i) // 保证栈中永远只有一个')'
			} else { // 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
				maxAns = max(maxAns, i-stack[len(stack)-1])
			}
		}
	}
	return maxAns
}

// 暴力
func longestValidParenthesesForce(s string) int {
	var isValid func(s string) bool
	isValid = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		stack := make([]rune, 0)
		for _, v := range s {
			if v == '(' { // 左括号入栈
				stack = append(stack, v)
			} else if v == ')' && len(stack) > 0 && stack[len(stack)-1] == '(' { // 右括号弹栈
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		return len(stack) == 0
	}

	if len(s) < 2 {
		return 0
	}
	n := len(s) / 2 * 2
	for i := n; i > 0; i = i - 2 {
		for j := 0; j < n-i+1; j++ {
			if isValid(s[j : j+i]) {
				return i
			}
		}
	}
	return 0
}
