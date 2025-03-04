package leetcode

// 小行星碰撞
// https://leetcode.cn/problems/asteroid-collision/

// 栈
func asteroidCollision(asteroids []int) []int {
	// 创建整数数组作为栈，正数元素入栈
	stack := make([]int, 0)

	for _, asteroid := range asteroids {
		// 当前行星是否存活
		alive := true

		// 当前行星向左移动(负数)且栈顶行星向右移动(正数)时会发生碰撞
		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			// 比较当前行星和栈顶行星的大小
			top := stack[len(stack)-1]

			if top < -asteroid {
				// 栈顶行星较小，被摧毁
				stack = stack[:len(stack)-1]
			} else if top == -asteroid {
				// 两个行星大小相等，都摧毁
				stack = stack[:len(stack)-1]
				alive = false
			} else {
				// 当前行星较小，被摧毁
				alive = false
			}
		}

		// 如果当前行星存活，加入栈中
		if alive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
