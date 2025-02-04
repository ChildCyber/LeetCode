package leetcode

// 最长有效括号
// https://leetcode-cn.com/problems/longest-valid-parentheses/

// 暴力
// 枚举一切可能子串，判断子串是不是有效的括号序列
func longestValidParenthesesBrute(s string) int {
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

	n := len(s) / 2 * 2            // 有效括号长度必然是偶数，取不超过 len(s) 的最大偶数作为起点
	for i := n; i > 0; i = i - 2 { // 从大到小枚举子串长度i
		for j := 0; j < n-i+1; j++ { // 从左到右枚举起点j，截取子串 s[j:j+i]
			if isValid(s[j : j+i]) {
				return i
			}
		}
	}
	return 0
}

// 栈
func longestValidParentheses(s string) int {
	// 用栈存下标，而不是存括号本身。压下标是因为要算长度，不能光知道匹配没匹配
	maxLen := 0
	stack := []int{-1} // 初始压一个-1作为基准。如果一开始栈为空，第一个字符为左括号的时候会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，在一开始的时候往栈中放入一个值为 −1 的元素

	for i := 0; i < len(s); i++ { // 遍历字符串
		if s[i] == '(' { // '('做入栈操作
			stack = append(stack, i) // 下标入栈
		} else { // ')'做出栈操作
			stack = stack[:len(stack)-1] // 先弹出栈顶元素，表示匹配了当前右括号
			if len(stack) == 0 {         // 如果栈为空，说明当前的右括号为没有被匹配的右括号，压入当前位置作为新的起点
				stack = append(stack, i) // 保证栈中永远只有一个')'
			} else { // 如果栈不为空，当前有效长度 = i - 栈顶
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 动态规划
func longestValidParenthesesDP(s string) int {
	// 状态定义：dp[i] 表示以 i 位置结尾的最长有效括号长度
	// 状态转移：
	// 如果 s[i] == ')' 并且 s[i-1] == '('，那就是 dp[i] = dp[i-2] + 2。
	// 如果 s[i] == ')' 并且 s[i-1] == ')'，就看 i - dp[i-1] - 1 位置是不是 '('，如果是：dp[i] = dp[i-1] + 2 + dp[i - dp[i-1] - 2]
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	maxLen := 0

	for i := 1; i < n; i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1]-2 >= 0 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
			if dp[i] > maxLen {
				maxLen = dp[i]
			}
		}
	}

	return maxLen
}

// 计数法（双向扫描）
func longestValidParenthesesCount(s string) int {
	// 从左到右：维护两个计数器 left 和 right，遇到 '(' 加 left，遇到 ')' 加 right。
	// 当 left == right → 更新答案。
	// 当 right > left → 清零（因为区间不可能再合法）
	// 从右到左：同样的逻辑，但角色互换，防止 (() 这种情况漏算
	maxLen := 0
	left, right := 0, 0

	// 从左往右
	for _, ch := range s {
		if ch == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left*2 > maxLen {
				maxLen = left * 2
			}
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	// 从右往左
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left*2 > maxLen {
				maxLen = left * 2
			}
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
