package leetcode

// 946. 验证栈序列
// https://leetcode.cn/problems/validate-stack-sequences/

// 借用一个辅助栈 stack ，模拟压入/弹出操作的排列。根据是否模拟成功，即可得到结果
// 入栈操作： 按照压栈序列的顺序执行。
// 出栈操作： 每次入栈后，循环判断 “栈顶元素 == 弹出序列的当前元素” 是否成立，将符合弹出序列顺序的栈顶元素全部弹出。
// 按照 push 数组的顺序先把压栈,然后再依次在 pop 里面找栈顶元素,找到了就
// pop,直到遍历完 pop 数组,最终如果遍历完了 pop 数组,就代表清空了整个栈了。
func validateStackSequences(pushed []int, popped []int) bool {
	stack, j, N := []int{}, 0, len(pushed)
	for _, x := range pushed {
		stack = append(stack, x) // 入栈
		// 循环出栈
		for len(stack) != 0 && j < N && stack[len(stack)-1] == popped[j] {
			stack = stack[0 : len(stack)-1] // 出栈
			j++
		}
	}
	return j == N
}
