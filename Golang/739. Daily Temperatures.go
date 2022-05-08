package leetcode

// 每日温度
// https://leetcode-cn.com/problems/daily-temperatures
// 暴力
func dailyTemperatures(temperatures []int) []int {
	ans, j := make([]int, len(temperatures)), 0
	for i := 0; i < len(temperatures); i++ {
		for j = i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] { // 找到第一个更大的元素
				ans[i] = j - i
				break
			}
		}
	}
	return ans
}

// 单调栈
func dailyTemperatures1(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{} // 单调栈

	// 遍历每日温度，维护一个存储下标的单调栈，从栈底到栈顶的下标对应的温度列表中的温度依次递减
	// 如果一个下标在单调栈里，则表示尚未找到下一次温度更高的下标
	for i := 0; i < length; i++ {
		temperature := temperatures[i]
		// 构建单调栈：当日温度大于栈顶温度，说明栈顶元素升温日已经找到，将栈顶元素出栈
		for len(stack) > 0 && temperature > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex
		}
		// 若栈为空或者当日温度小于等于栈顶温度则直接入栈
		stack = append(stack, i) // 存储下标
	}
	return ans
}
