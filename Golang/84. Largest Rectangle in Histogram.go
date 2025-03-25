package leetcode

// 柱状图中最大的矩形
// https://leetcode.cn/problems/largest-rectangle-in-histogram/

// 问题本质：矩形高度由最矮的柱子决定。对每个柱子，找到左右第一个比它矮的柱子

// 暴力
// 检查所有可能的矩形区域，找出面积最大的那个
// 时间复杂度：O(n³)
func largestRectangleAreaBrute(heights []int) int {
	// 思路：枚举所有区间+取最小高度
	// 固定左右边界
	// 1. 遍历所有可能的左边界（从第一个柱子到最后一个柱子）
	// 2. 遍历所有可能的右边界（从左边界到最后一个柱子）
	// 3. 对于每个区间 [左边界, 右边界]：
	// 找到这个区间内最矮的柱子高度
	// 计算矩形面积 = 最矮高度 × (右边界 - 左边界 + 1)
	// 4. 记录遇到的最大面积
	n := len(heights)
	maxArea := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// 找 i..j 的最小高度
			minH := heights[i]
			for k := i; k <= j; k++ {
				if heights[k] < minH {
					minH = heights[k]
				}
			}
			area := minH * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

// 暴力优化
// 时间复杂度：O(n²)
func largestRectangleAreaBruteOptimized(heights []int) int {
	// 思路：枚举柱子作为最矮
	// 固定高度，扩展宽度
	// 1. 以每个柱子作为矩形的高度
	// 2. 向左右两侧扩展，直到遇到更矮的柱子
	// 3. 计算以当前柱子为高度的最大矩形面积
	// 4. 记录遇到的最大面积
	n := len(heights)
	maxArea := 0
	for i := 0; i < n; i++ {
		height := heights[i]

		// 向左扩展
		left := i
		for left > 0 && heights[left-1] >= height {
			left--
		}

		// 向右扩展
		right := i
		for right < n-1 && heights[right+1] >= height {
			right++
		}

		area := height * (right - left + 1)
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}

// 单调栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func largestRectangleArea(heights []int) int {
	/* 思路：
	   1. 用一个栈存储柱子的下标，保持栈里的高度递增
	   2. 遍历过程
	   对于每个柱子i：
	   1) 如果栈为空或当前柱子高度 ≥ 栈顶柱子高度：
	     将当前柱子索引压入栈中
	   2) 如果当前柱子高度 < 栈顶柱子高度：
	     不断弹出栈顶元素，直到栈为空或栈顶柱子高度 ≤ 当前柱子高度
	     对于每个弹出的栈顶元素top：
	       高度 = heights[top]
	       宽度 = i - 栈顶元素索引 - 1（如果栈为空，宽度 = i）
	       面积 = 高度 × 宽度
	       更新最大面积
	   3. 处理剩余栈中元素
	   遍历完成后，如果栈中还有元素：
	     依次弹出并计算每个柱子能形成的矩形面积
	     宽度 = 总柱子数 - 栈顶元素索引 - 1（如果栈为空，宽度 = 总柱子数）
	   单调栈有效的原因：栈中柱子高度递增，当遇到较矮柱子时，可以确定之前较高柱子的右边界，再通过栈记录潜在的"左边界"
	   矩形的面积由最矮柱子的高度和能延伸的宽度决定，单调栈帮助快速找到每个柱子作为最矮柱子时的左右边界，当遇到较矮柱子时，之前较高的柱子不能再向右延伸，此时计算它们的最大面积
	*/

	// 创建一个栈，存储柱子的索引
	stack := make([]int, 0)
	maxArea := 0

	// 补一个0在末尾，方便清算。如果数组走到末尾，一直没有遇到比最后一根更矮的，栈里就会残留一堆没结算过的高度
	heights = append(heights, 0)

	// 遍历所有柱子
	for i := 0; i < len(heights); i++ {
		// 当栈不为空且当前柱子高度小于栈顶柱子高度时
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// 弹出栈顶元素
			poppedIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 计算宽度
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}

			// 计算面积并更新最大值
			area := heights[poppedIndex] * width
			if area > maxArea {
				maxArea = area
			}
		}

		// 将当前柱子索引压入栈中
		stack = append(stack, i)
	}

	return maxArea
}
